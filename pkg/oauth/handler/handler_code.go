package handler

import (
	"context"
	"github.com/imulab/soteria/pkg/oauth"
	"github.com/imulab/soteria/pkg/oauth/client"
	oauthError "github.com/imulab/soteria/pkg/oauth/error"
	"github.com/imulab/soteria/pkg/oauth/request"
	"github.com/imulab/soteria/pkg/oauth/token"
	"github.com/imulab/soteria/pkg/utility"
)

type AuthorizeCodeHandler struct {
	ScopeStrategy 	client.ScopeStrategy
	CodeStrategy 	token.AuthorizeCodeStrategy
	CodeStorage 	token.AuthorizeCodeRepository
	next 			OAuthAuthorizeHandler
}

func (h *AuthorizeCodeHandler) Next() OAuthAuthorizeHandler {
	return h.next
}

func (h *AuthorizeCodeHandler) WithNext(next OAuthAuthorizeHandler) OAuthAuthorizeHandler {
	h.next = next
	return h.next
}

func (h *AuthorizeCodeHandler) Authorize(req request.OAuthAuthorizeRequest, resp oauth.AuthorizeResponse, ctx context.Context) error {
	select {
	case <-ctx.Done():
		return oauthError.ContextCancelled()
	default:
		// continue processing
	}

	if !utility.Exactly(req.GetClient().GetResponseTypes(), oauth.ResponseTypeCode) {
		return h.nextAuthorize(req, resp, ctx)
	}

	if !h.ScopeStrategy.AcceptsAll(req.GetClient(), req.GetSession().GetGrantedScopes()) {
		return oauthError.InvalidScope("rejected by client.")
	}

	if code, err := h.CodeStrategy.NewCode(req, ctx); err != nil {
		return err
	} else if err := h.CodeStorage.Save(code, req, ctx); err != nil {
		return err
	} else {
		resp.SetCode(code)
	}

	req.HandledResponseType(oauth.ResponseTypeCode)

	return h.nextAuthorize(req, resp, ctx)
}

func (h *AuthorizeCodeHandler) nextAuthorize(req request.OAuthAuthorizeRequest, resp oauth.AuthorizeResponse, ctx context.Context) error {
	if h.next != nil {
		return h.next.Authorize(req, resp, ctx)
	}
	return nil
}

