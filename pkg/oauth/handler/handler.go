package handler

import (
	"context"
	"github.com/imulab/soteria/pkg/oauth"
	"github.com/imulab/soteria/pkg/oauth/request"
)

// Interface for handling an OAuth 2.0 authorize request.
type OAuthAuthorizeHandler interface {
	// Handle the authorization request.
	Authorize(req request.OAuthAuthorizeRequest, resp oauth.AuthorizeResponse, ctx context.Context) error
	// Returns the next handler in chain
	Next() OAuthAuthorizeHandler
	// Sets the next handler and returns the newly set handler (for fluency)
	WithNext(next OAuthAuthorizeHandler) OAuthAuthorizeHandler
}

// Interface for handling an OAuth 2.0 token request.
type OAuthTokenHandler interface {
	// Update session knowledge before processing request.
	UpdateSession(req request.OAuthTokenRequest, ctx context.Context) error
	// Issue token for the request.
	IssueAccess(req request.OAuthTokenRequest, resp oauth.TokenResponse, ctx context.Context) error
	// Returns next handler in chain
	Next() OAuthTokenHandler
	// Sets the next handler and returns the newly set handler (for fluency)
	WithNext(next OAuthTokenHandler) OAuthTokenHandler
}