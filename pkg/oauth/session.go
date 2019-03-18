package oauth

import "github.com/imulab/soteria/pkg/utility"

// Interface for a user session during the request
type Session interface {
	// Returns the user subject that authorized this request.
	GetSubject() string
	// Returns user's granted scopes.
	GetGrantedScopes() []string
	// Returns claims to be added in the issued access token.
	GetAccessClaims() map[string]interface{}
	// Clone the user session
	Clone() Session
}

func NewDefaultSession() Session {
	return &DefaultSession{
		Scopes: make([]string, 0),
		Claims: make(map[string]interface{}),
	}
}

// Implementation for a default user session
type DefaultSession struct {
	Subject 	string					`json:"subject"`
	Scopes		[]string				`json:"scopes"`
	Claims 		map[string]interface{}	`json:"claims"`
}

func (s *DefaultSession) GetSubject() string {
	return s.Subject
}

func (s *DefaultSession) GetGrantedScopes() []string {
	return s.Scopes
}

func (s *DefaultSession) GetAccessClaims() map[string]interface{} {
	return s.Claims
}

func (s *DefaultSession) Clone() Session {
	return &DefaultSession{
		Subject: s.Subject,
		Scopes:  utility.CopyStrArr(s.Scopes),
		Claims:  utility.CopyStringGenericMap(s.Claims),
	}
}
