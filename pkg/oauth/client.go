package oauth

// Interface to an OAuth 2.0 Client.
type Client interface {
	// Returns an universal client identifier.
	GetId() string
	// Returns the client's name.
	GetName() string
	// Returns client type.
	// [confidential|public]
	GetType() string
	// Returns registered redirect URIs.
	GetRedirectUris() []string
	// Returns registered response types.
	// [code|token]
	GetResponseTypes() []string
	// Returns registered grant types.
	// [authorization_code|implicit|password|client_credentials|refresh_token]
	GetGrantTypes() []string
	// Returns registered scopes.
	GetScopes() []string
}

// Default implementation to Client.
type DefaultClient struct {
	Id 				string		`json:"client_id"`
	Name 			string		`json:"client_name"`
	Type 			string		`json:"client_type"`
	RedirectUris	[]string	`json:"redirect_uris"`
	ResponseTypes 	[]string	`json:"response_types"`
	GrantTypes 		[]string	`json:"grant_types"`
	Scopes 			[]string	`json:"scopes"`
}

func (c *DefaultClient) GetId() string {
	return c.Id
}

func (c *DefaultClient) GetName() string {
	return c.Name
}

func (c *DefaultClient) GetType() string {
	return c.Type
}

func (c *DefaultClient) GetRedirectUris() []string {
	return c.RedirectUris
}

func (c *DefaultClient) GetResponseTypes() []string {
	return c.ResponseTypes
}

func (c *DefaultClient) GetGrantTypes() []string {
	return c.GrantTypes
}

func (c *DefaultClient) GetScopes() []string {
	return c.Scopes
}

// Single purpose interface for finding a client by its id.
type ClientLookup interface {
	// Find client by its id. Returns the client or an ErrClientNotFound error.
	Find(id string) (Client, error)
}

// Dummy implementation for ClientLookup that always returns not found error.
type NotFoundClientLookup struct {}

func (_ *NotFoundClientLookup) Find(id string) (Client, error) {
	return nil, ErrClientNotFound
}
