package error

import "fmt"

type OAuthError struct {
	Err 		string	`json:"error"`
	Description	string	`json:"error_description"`
}

func (e *OAuthError) Error() string {
	return fmt.Sprintf("%s:%s", e.Err, e.Description)
}

// invalid_request
// The request is missing a required parameter, includes an
// unsupported parameter value (other than grant type),
// repeats a parameter, includes multiple credentials,
// utilizes more than one mechanism for authenticating the
// client, or is otherwise malformed.
func InvalidRequest(description string) error {
	return &OAuthError{
		Err: "invalid_request",
		Description: description,
	}
}

// invalid_client
// Client authentication failed (e.g., unknown client, no
// client authentication included, or unsupported
// authentication method).  The authorization server MAY
// return an HTTP 401 (Unauthorized) status code to indicate
// which HTTP authentication schemes are supported.  If the
// client attempted to authenticate via the "Authorization"
// request header field, the authorization server MUST
// respond with an HTTP 401 (Unauthorized) status code and
// include the "WWW-Authenticate" response header field
// matching the authentication scheme used by the client.
func InvalidClient(description string) error {
	return &OAuthError{
		Err: "invalid_client",
		Description: description,
	}
}

// invalid_grant
// The provided authorization grant (e.g., authorization
// code, resource owner credentials) or refresh token is
// invalid, expired, revoked, does not match the redirection
// URI used in the authorization request, or was issued to
// another client.
func InvalidGrant(description string) error {
	return &OAuthError{
		Err: "invalid_grant",
		Description: description,
	}
}

// unauthorized_client
// The authenticated client is not authorized to use this
// authorization grant type.
func UnauthorizedClient(description string) error {
	return &OAuthError{
		Err: "unauthorized_client",
		Description: description,
	}
}

// unsupported_grant_type
// The authorization grant type is not supported by the
// authorization server.
func UnsupportedGrantType(description string) error {
	return &OAuthError{
		Err: "unsupported_grant_type",
		Description: description,
	}
}

// unsupported_response_type
// The authorization server does not support obtaining an
// authorization code using this method.
func UnsupportedResponseType(description string) error {
	return &OAuthError{
		Err: "unsupported_response_type",
		Description: description,
	}
}

// invalid_scope
// The requested scope is invalid, unknown, malformed, or
// exceeds the scope granted by the resource owner.
func InvalidScope(description string) error {
	return &OAuthError{
		Err: "invalid_scope",
		Description: description,
	}
}

// access_denied
// The resource owner or authorization server denied the
// request.
func AccessDenied(description string) error {
	return &OAuthError{
		Err: "access_denied",
		Description: description,
	}
}

// server_error
// The authorization server encountered an unexpected
// condition that prevented it from fulfilling the request.
// (This error code is needed because a 500 Internal Server
// Error HTTP status code cannot be returned to the client
// via an HTTP redirect.)
func ServerError(description string) error {
	return &OAuthError{
		Err: "server_error",
		Description: description,
	}
}

// temporarily_unavailable
// The authorization server is currently unable to handle
// the request due to a temporary overloading or maintenance
// of the server.  (This error code is needed because a 503
// Service Unavailable HTTP status code cannot be returned
// to the client via an HTTP redirect.)
func TemporarilyUnavailable(description string) error {
	return &OAuthError{
		Err: "temporarily_unavailable",
		Description: description,
	}
}
