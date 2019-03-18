package request

type OAuthTokenRequest interface {
	OAuthRequest
	// Get the grant types
	GetGrantTypes() []string
	// Get the supplied authorization code
	GetCode() string
	// Get the supplied refresh token
	GetRefreshToken() string
}
