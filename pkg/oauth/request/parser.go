package request

import (
	"context"
	"net/http"
)

type OAuthAuthorizeRequestParser interface {
	// Returns the next parser in chain
	Next() OAuthAuthorizeRequestParser
	// Convenience method to chain the next parser, returns the next parser.
	WithNext(next OAuthAuthorizeRequestParser) OAuthAuthorizeRequestParser
	// Parse HTTP request (and/or other data) into the OAuthAuthorizeRequestBuilder.
	Parse(ctx context.Context, r *http.Request, req OAuthAuthorizeRequest) error
}