package request

import (
	"github.com/imulab/soteria/pkg/oauth"
	"github.com/imulab/soteria/pkg/oauth/client"
	"time"
)

// Interface for shared elements in all OAuth 2.0 requests.
type OAuthRequest interface {
	// Returns the id of the request
	GetId()	string
	// Set the id
	setId(id string)
	// Returns the request timestamp
	GetTimestamp() time.Time
	// Set timestamp
	setTimestamp(timestamp int64)
	// Returns the client for this request
	GetClient() client.Client
	// Set client
	setClient(client client.Client)
	// Returns the requested redirect uri
	GetRedirectUri() string
	// Set redirect uri
	setRedirectUri(uri string)
	// Returns the session
	GetSession() oauth.Session
	// Set session
	setSession(session oauth.Session)
}

// Implementation for OAuthRequest
type oauthRequestImpl struct {
	Id 			string
	Timestamp	int64
	Client 		client.Client
	RedirectUri	string
	Session 	oauth.Session
}

func (r *oauthRequestImpl) setId(id string) {
	r.Id = id
}

func (r *oauthRequestImpl) setTimestamp(timestamp int64) {
	r.Timestamp = timestamp
}

func (r *oauthRequestImpl) setClient(client client.Client) {
	r.Client = client
}

func (r *oauthRequestImpl) setRedirectUri(uri string) {
	r.RedirectUri = uri
}

func (r *oauthRequestImpl) setSession(session oauth.Session) {
	r.Session = session
}

func (r *oauthRequestImpl) GetId() string {
	return r.Id
}

func (r *oauthRequestImpl) GetTimestamp() time.Time {
	return time.Unix(r.Timestamp, 0)
}

func (r *oauthRequestImpl) GetClient() client.Client {
	return r.Client
}

func (r *oauthRequestImpl) GetRedirectUri() string {
	return r.RedirectUri
}

func (r *oauthRequestImpl) GetSession() oauth.Session {
	return r.Session
}
