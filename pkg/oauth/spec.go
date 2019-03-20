package oauth

// Response Types
const (
	ResponseTypeCode = "code"
	ResponseTypeToken = "token"
)

// Grant Types
const (
	GrantTypeCode = "authorization_code"
	GrantTypeImplicit = "implicit"
	GrantTypePassword = "password"
	GrantTypeClient = "client_credentials"
	GrantTypeRefresh = "refresh_token"
)

// Standard scopes
const (
	ScopeOffline = "offline"
	ScopeOfflineAccess = "offline_access"
)

// Parameters
const (
	ParamClientId = "client_id"
	ParamResponseType = "response_type"
	ParamRedirectUri = "redirect_uri"
	ParamScope = "scope"
	ParamState = "state"
)