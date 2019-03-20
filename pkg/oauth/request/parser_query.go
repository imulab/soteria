package request

import (
	"context"
	"fmt"
	"github.com/imulab/soteria/pkg/oauth"
	"github.com/imulab/soteria/pkg/oauth/client"
	oauthError "github.com/imulab/soteria/pkg/oauth/error"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type OAuthAuthorizeRequestQueryParser struct {
	next                       OAuthAuthorizeRequestParser
	ClientLookup               client.Repository
	ClientLookupTimeoutSeconds uint8
}

func (p *OAuthAuthorizeRequestQueryParser) clientLookupTimeout() time.Duration {
	sec := p.ClientLookupTimeoutSeconds
	if sec == 0 {
		sec = 5
	}
	return time.Duration(sec) * time.Second
}

func (p *OAuthAuthorizeRequestQueryParser) Next() OAuthAuthorizeRequestParser {
	return p.next
}

func (p *OAuthAuthorizeRequestQueryParser) WithNext(next OAuthAuthorizeRequestParser) OAuthAuthorizeRequestParser {
	p.next = next
	return p.next
}

func (p *OAuthAuthorizeRequestQueryParser) Parse(ctx context.Context, r *http.Request, req OAuthAuthorizeRequest) error {
	switch r.Method {
	case http.MethodGet:
		if queries, err := url.ParseQuery(r.URL.RawQuery); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Errorln("failed to parse url query.")
			return oauthError.ServerError("failed to parse url query.")
		} else if err := p.parseHttpGet(ctx, queries, req); err != nil {
			return errors.WithStack(err)
		}
	default:
		return oauthError.InvalidRequest("unsupported method.")
	}

	if p.next != nil {
		return p.next.Parse(ctx, r, req)
	}
	return nil
}

func (p *OAuthAuthorizeRequestQueryParser) parseHttpGet(ctx context.Context, v url.Values, req OAuthAuthorizeRequest) error {
	select {
	case <-ctx.Done():
		return oauthError.ContextCancelled()
	default:
		// continue
	}

	// debug
	logrus.WithFields(logrus.Fields{
		oauth.ParamClientId: v.Get(oauth.ParamClientId),
		oauth.ParamResponseType: v.Get(oauth.ParamResponseType),
		oauth.ParamRedirectUri: v.Get(oauth.ParamRedirectUri),
		oauth.ParamScope: v.Get(oauth.ParamScope),
		oauth.ParamState: v.Get(oauth.ParamState),
	}).Debug("received request.")

	// client
	var c client.Client
	findClientChan, findClientErr := make(chan client.Client), make(chan error)
	findClientCtx, cancelFindClient := context.WithTimeout(ctx, p.clientLookupTimeout())
	defer cancelFindClient()
	go p.findClient(findClientCtx, v.Get(oauth.ParamClientId), findClientChan, findClientErr)

	// client independent parameters
	req.addResponseTypes(strings.Split(v.Get(oauth.ParamResponseType), " "))
	req.addScopes(strings.Split(v.Get(oauth.ParamScope), " "))
	req.setState(v.Get(oauth.ParamState))

	// wait for client result
	select {
	case <-findClientCtx.Done():
		return oauthError.ContextCancelled()
	case err := <-findClientErr:
		return err
	case c = <-findClientChan:
		req.setClient(c)
	}

	// redirect_uri
	if effectiveRedirectUri, err := oauth.SelectRedirectUri(v.Get(oauth.ParamRedirectUri), c.GetRedirectUris()); err != nil {
		return errors.WithStack(err)
	} else {
		req.setRedirectUri(effectiveRedirectUri)
	}

	return nil
}

func (p *OAuthAuthorizeRequestQueryParser) findClient(ctx context.Context, clientId string, resultChan chan <-client.Client, errChan chan <-error) {
	if len(clientId) == 0 {
		errChan <- oauthError.InvalidRequest(fmt.Sprintf("%s is required.", oauth.ParamClientId))
	} else if c, err := p.ClientLookup.Find(ctx, clientId); err != nil {
		errChan <- errors.WithStack(err)
	} else {
		resultChan <- c
	}
}