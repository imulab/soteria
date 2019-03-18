package token

import (
	"context"
	"github.com/imulab/soteria/pkg/oauth"
	"github.com/imulab/soteria/pkg/oauth/request"
)

// Interface for creating, verifying and revoking an access token
type AccessTokenStrategy interface {
	// Generate a new access token.
	NewToken(req request.OAuthRequest, ctx context.Context) (string, error)
	// Validate the given access token.
	ValidateToken(token string, req request.OAuthTokenRequest, ctx context.Context) error
}

// Interface for managing persistence of access token
type AccessTokenRepository interface {
	// Find the associated session with this token
	GetSession(token string, ctx context.Context) (oauth.Session, error)
	// Persist the access token and associate it with the request session
	Save(token string, req request.OAuthRequest, ctx context.Context) error
	// Remove the access token from persistence
	Delete(token string, ctx context.Context) error
}