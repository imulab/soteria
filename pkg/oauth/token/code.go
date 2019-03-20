package token

import (
	"context"
	"github.com/imulab/soteria/pkg/crypt"
	"github.com/imulab/soteria/pkg/oauth"
	oauthError "github.com/imulab/soteria/pkg/oauth/error"
	"github.com/imulab/soteria/pkg/oauth/request"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

// Interface for creating and verifying an authorization code
type AuthorizeCodeStrategy interface {
	// Generate a new code.
	NewCode(req request.OAuthAuthorizeRequest, ctx context.Context) (string, error)
	// Validate the given authorize code, update session information, if necessary
	ValidateCode(code string, req request.OAuthAuthorizeRequest, ctx context.Context) error
}

// Factory method to construct an AuthorizeCodeStrategy using HMAC-SHA256
func NewHmacSha256AuthorizeCodeStrategy(entropy uint, signingKey []byte) (AuthorizeCodeStrategy, error) {
	strategy, err := crypt.NewHmacSha256Strategy(signingKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &HmacShaAuthorizeCodeStrategy{
		Entropy: entropy,
		Hmac: strategy,
	}, nil
}

// Factory method to construct an AuthorizeCodeStrategy using HMAC-SHA384
func NewHmacSha384AuthorizeCodeStrategy(entropy uint, signingKey []byte) (AuthorizeCodeStrategy, error) {
	strategy, err := crypt.NewHmacSha384Strategy(signingKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &HmacShaAuthorizeCodeStrategy{
		Entropy: entropy,
		Hmac: strategy,
	}, nil
}

// Factory method to construct an AuthorizeCodeStrategy using HMAC-SHA512
func NewHmacSha512AuthorizeCodeStrategy(entropy uint, signingKey []byte) (AuthorizeCodeStrategy, error) {
	strategy, err := crypt.NewHmacSha512Strategy(signingKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &HmacShaAuthorizeCodeStrategy{
		Entropy: entropy,
		Hmac: strategy,
	}, nil
}

// Implementation of AuthorizeCodeStrategy using HMAC-SHA series algorithms
type HmacShaAuthorizeCodeStrategy struct {
	// Length of bytes for the key part of the generated code
	Entropy	uint
	// Hmac-Sha strategy
	Hmac	crypt.HmacShaStrategy
}

func (s *HmacShaAuthorizeCodeStrategy) NewCode(req request.OAuthAuthorizeRequest, ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", oauthError.ContextCancelled()
	default:
		if key, sig, err := s.Hmac.Generate(s.Entropy); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Errorln("failed to generate authorization code.")
			return "", oauthError.ServerError("failed to generate authorization code.")
		} else {
			return key + "." + sig, nil
		}
	}
}

func (s *HmacShaAuthorizeCodeStrategy) ValidateCode(code string, req request.OAuthAuthorizeRequest, ctx context.Context) error {
	select {
	case <-ctx.Done():
		return oauthError.ContextCancelled()
	default:
		parts := strings.Split(code, ".")
		if len(parts) != 2 {
			logrus.Debugln("authorization code has bad format.")
			return oauthError.InvalidGrant("invalid authorization code.")
		} else if err := s.Hmac.Verify(parts[0], parts[1]); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Debugln("authorization code has bad signature.")
			return oauthError.InvalidGrant("invalid authorization code.")
		}
		return nil
	}
}

type AuthorizeCodeRepository interface {
	// Find the associated session with the authorization code
	GetSession(code string, ctx context.Context) (oauth.Session, error)
	// Interface for managing persistence of authorization code
	// Persist the authorize code and associate it with the request session
	Save(code string, req request.OAuthAuthorizeRequest, ctx context.Context) error
	// Remove the authorize code from persistence
	Delete(code string, ctx context.Context) error
}

// Factory method for creating an in memory AuthorizeCodeRepository
func NewMemoryAuthorizeCodeRepository() AuthorizeCodeRepository {
	return &memoryAuthorizeCodeRepository{db:make(map[string]oauth.Session)}
}

// In memory implementation for AuthorizeCodeRepository. Not intended for produce usage.
type memoryAuthorizeCodeRepository struct {
	sync.RWMutex
	db 	map[string]oauth.Session
}

func (r *memoryAuthorizeCodeRepository) GetSession(code string, ctx context.Context) (oauth.Session, error) {
	select {
	case <-ctx.Done():
		return nil, oauthError.ContextCancelled()
	default:
		r.RLock()
		defer r.RUnlock()
		if session, ok := r.db[code]; !ok {
			return nil, oauthError.InvalidGrant("authorization code not found.")
		} else {
			return session, nil
		}
	}
}

func (r *memoryAuthorizeCodeRepository) Save(code string, req request.OAuthAuthorizeRequest, ctx context.Context) error {
	select {
	case <-ctx.Done():
		return oauthError.ContextCancelled()
	default:
		r.Lock()
		defer r.Unlock()
		r.db[code] = req.GetSession()
		return nil
	}
}

func (r *memoryAuthorizeCodeRepository) Delete(code string, ctx context.Context) error {
	select {
	case <-ctx.Done():
		return oauthError.ContextCancelled()
	default:
		r.Lock()
		defer r.Unlock()
		delete(r.db, code)
		return nil
	}
}

func (r *memoryAuthorizeCodeRepository) Clear() {
	r.Lock()
	defer r.Unlock()
	r.db = make(map[string]oauth.Session)
}

var (
	noOpAuthorizeCodeRepositoryInstance = &noOpAuthorizeCodeRepository{}
)

// Convenience method for accessing the no-op implementation of AuthorizeCodeRepository
func NewNoOpAuthorizeCodeRepository() AuthorizeCodeRepository {
	return noOpAuthorizeCodeRepositoryInstance
}

// No-op implementation for AuthorizeCodeRepository. This implementation does not support the GetSession method and
// will panic once called. All other methods is implemented as no-op.
type noOpAuthorizeCodeRepository struct {}

func (_ *noOpAuthorizeCodeRepository) GetSession(code string, ctx context.Context) (oauth.Session, error) {
	panic("should not be called")
}

func (_ *noOpAuthorizeCodeRepository) Save(code string, req request.OAuthAuthorizeRequest, ctx context.Context) error {
	return nil
}

func (_ *noOpAuthorizeCodeRepository) Delete(code string, ctx context.Context) error {
	return nil
}
