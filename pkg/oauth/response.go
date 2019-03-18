package oauth

// Interface for OAuth 2.0 authorize response.
type AuthorizeResponse interface {
	// Returns the assigned authorization code
	GetCode() string
	// Assign the authorization code
	SetCode(code string)
	// Returns the confirmed redirection uri
	GetRedirectUri() string
	// Set the confirmed redirection uri
	SetRedirectUri(uri string)
}

func NewDefaultAuthorizeResponse() AuthorizeResponse {
	return &DefaultAuthorizeResponse{}
}

// Default implementation of AuthorizeResponse
type DefaultAuthorizeResponse struct {
	code 		string
	redirectUri	string
}

func (r *DefaultAuthorizeResponse) GetCode() string {
	return r.code
}

func (r *DefaultAuthorizeResponse) SetCode(code string) {
	r.code = code
}

func (r *DefaultAuthorizeResponse) GetRedirectUri() string {
	return r.redirectUri
}

func (r *DefaultAuthorizeResponse) SetRedirectUri(uri string) {
	r.redirectUri = uri
}

// Interface for OAuth 2.0 token response.
type TokenResponse interface {
	// Returns the assigned access token
	GetAccessToken() string
	// Sets the assigned access token
	SetAccessToken(token string)
	// Returns the assigned refresh token
	GetRefreshToken() string
	// Sets the assigned refresh token
	SetRefreshToken(token string)
}

// Default implementation for the TokenResponse.
type DefaultTokenResponse struct {
	accessToken 	string
	refreshToken 	string
}

func (r *DefaultTokenResponse) GetAccessToken() string {
	return r.accessToken
}

func (r *DefaultTokenResponse) SetAccessToken(token string) {
	r.accessToken = token
}

func (r *DefaultTokenResponse) GetRefreshToken() string {
	return r.refreshToken
}

func (r *DefaultTokenResponse) SetRefreshToken(token string) {
	r.refreshToken = token
}
