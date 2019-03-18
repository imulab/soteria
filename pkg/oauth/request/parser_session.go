package request

import (
	"context"
	"github.com/imulab/soteria/pkg/oauth"
	"net/http"
)

type OAuthAuthorizeRequestSessionParser struct {
	next 		OAuthAuthorizeRequestParser
}

func (p *OAuthAuthorizeRequestSessionParser) Next() OAuthAuthorizeRequestParser {
	return p.next
}

func (p *OAuthAuthorizeRequestSessionParser) WithNext(next OAuthAuthorizeRequestParser) OAuthAuthorizeRequestParser {
	p.next = next
	return p.next
}

func (p *OAuthAuthorizeRequestSessionParser) Parse(ctx context.Context, r *http.Request, req OAuthAuthorizeRequest) error {
	req.setSession(oauth.NewDefaultSession())
	return nil
}

