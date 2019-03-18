package oauth

import "errors"

var (
	ErrContextCancelled = errors.New("context was cancelled")
)

var (
	ErrSignatureMismatch = errors.New("signature mismatch")
	ErrInvalidAuthorizeCode = errors.New("invalid authorize code")
	ErrAuthorizeCodeNotFound = errors.New("authorize code not found")
	ErrClientNotFound = errors.New("client is not found")
	ErrMethodNotSupported = errors.New("http method not supported")
	ErrMissingParam = errors.New("missing parameter")

	// redirect_uri
	ErrMultipleRedirectUri = errors.New("multiple registered redirect_uri")
	ErrUnregisteredRedirectUri = errors.New("redirect_uri not registered")
)