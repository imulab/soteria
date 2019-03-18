package token

import (
	"context"
	"github.com/imulab/soteria/pkg/oauth/request"
)

// Interface for creating, verifying and revoking refresh token
type RefreshTokenStrategy interface {
	// Generate a new refresh token.
	NewToken(req request.OAuthTokenRequest, ctx context.Context) (string, error)
	// Validate the given refresh token.
	ValidateToken(token string, req request.OAuthTokenRequest, ctx context.Context) error
}

// Interface for managing persistence of refresh token
type RefreshTokenRepository interface {
	// Persist the refresh token and associate it with the request session
	Save(token string, req request.OAuthTokenRequest, ctx context.Context) error
	// Remove the refresh token from persistence
	Delete(token string, ctx context.Context) error
}