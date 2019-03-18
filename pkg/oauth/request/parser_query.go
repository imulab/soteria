package request

import (
	"context"
	"github.com/imulab/soteria/pkg/oauth"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type OAuthAuthorizeRequestQueryParser struct {
	next                       OAuthAuthorizeRequestParser
	ClientLookup               oauth.ClientLookup
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
			return errors.WithStack(err)
		} else if err := p.parseHttpGet(ctx, queries, req); err != nil {
			return errors.WithStack(err)
		}
	default:
		return oauth.ErrMethodNotSupported
	}

	if p.next != nil {
		return p.next.Parse(ctx, r, req)
	}
	return nil
}

func (p *OAuthAuthorizeRequestQueryParser) parseHttpGet(ctx context.Context, v url.Values, req OAuthAuthorizeRequest) error {
	select {
	case <-ctx.Done():
		return oauth.ErrContextCancelled
	default:
		// continue
	}

	// debug
	logrus.WithFields(logrus.Fields{
		"client_id": v.Get("client_id"),
		"response_type": v.Get("response_type"),
		"redirect_uri": v.Get("redirect_uri"),
		"scope": v.Get("scope"),
		"state": v.Get("state"),
	}).Debug("Received request.")

	// client
	var client oauth.Client
	findClientChan, findClientErr := make(chan oauth.Client), make(chan error)
	findClientCtx, cancelFindClient := context.WithTimeout(ctx, p.clientLookupTimeout())
	defer cancelFindClient()
	go p.findClient(findClientCtx, v.Get("client_id"), findClientChan, findClientErr)

	// client independent parameters
	req.addResponseTypes(strings.Split(v.Get("response_type"), " "))
	req.addScopes(strings.Split(v.Get("scope"), " "))
	req.setState(v.Get("state"))

	// wait for client result
	select {
	case <-findClientCtx.Done():
		return oauth.ErrContextCancelled
	case err := <-findClientErr:
		return err
	case client = <-findClientChan:
		req.setClient(client)
	}

	// redirect_uri
	if effectiveRedirectUri, err := oauth.SelectRedirectUri(v.Get("redirect_uri"), client.GetRedirectUris()); err != nil {
		return errors.WithStack(err)
	} else {
		req.setRedirectUri(effectiveRedirectUri)
	}

	return nil
}

func (p *OAuthAuthorizeRequestQueryParser) findClient(ctx context.Context, clientId string, resultChan chan <-oauth.Client, errChan chan <-error) {
	if len(clientId) == 0 {
		errChan <- oauth.ErrMissingParam
	} else if client, err := p.ClientLookup.Find(clientId); err != nil {
		errChan <- errors.WithStack(err)
	} else {
		resultChan <- client
	}
}