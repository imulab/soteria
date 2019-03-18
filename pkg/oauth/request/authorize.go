package request

import (
	"github.com/satori/go.uuid"
	"time"
)

func NewOAuthAuthorizeRequest() OAuthAuthorizeRequest {
	return &oauthAuthorizeRequestImpl {
		oauthRequestImpl: &oauthRequestImpl{
			Id: uuid.NewV4().String(),
			Timestamp: time.Now().Unix(),
		},
		handleMap: make(map[string]struct{}),
	}
}

// Interface for an OAuth 2.0 authorize request
type OAuthAuthorizeRequest interface {
	OAuthRequest
	// Get requested response types
	GetResponseTypes() []string
	// Add response types
	addResponseTypes(responseTypes []string)
	// Get requested scopes
	GetScopes() []string
	// Add scopes
	addScopes(scopes []string)
	// Get the supplied state parameter
	GetState() string
	// Set state
	setState(state string)
	// Set the response type to handled
	HandledResponseType(responseType string)
	// Returns true if the response type has been handled; false otherwise
	IsResponseTypeHandled(responseType string) bool
}

// Implementation for OAuthAuthorizeRequest
type oauthAuthorizeRequestImpl struct {
	*oauthRequestImpl
	responseTypes []string
	scopes        []string
	state         string
	handleMap     map[string]struct{}
}

func (r *oauthAuthorizeRequestImpl) addResponseTypes(responseTypes []string) {
	if r.responseTypes == nil {
		r.responseTypes = make([]string, 0)
	}
	r.responseTypes = append(r.responseTypes, responseTypes...)
}

func (r *oauthAuthorizeRequestImpl) addScopes(scopes []string) {
	if r.scopes == nil {
		r.scopes = make([]string, 0)
	}
	r.scopes = append(r.scopes, scopes...)
}

func (r *oauthAuthorizeRequestImpl) setState(state string) {
	r.state = state
}

func (r *oauthAuthorizeRequestImpl) GetResponseTypes() []string {
	if r.responseTypes == nil {
		return []string{}
	}
	return r.responseTypes
}

func (r *oauthAuthorizeRequestImpl) GetScopes() []string {
	if r.scopes == nil {
		return []string{}
	}
	return r.scopes
}

func (r *oauthAuthorizeRequestImpl) GetState() string {
	return r.state
}

func (r *oauthAuthorizeRequestImpl) HandledResponseType(responseType string) {
	r.handleMap[responseType] = struct{}{}
}

func (r *oauthAuthorizeRequestImpl) IsResponseTypeHandled(responseType string) bool {
	_, ok := r.handleMap[responseType]
	return ok
}
