package handler

import (
	"context"
	"github.com/imulab/soteria/pkg/oauth"
	"github.com/imulab/soteria/pkg/oauth/client"
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
		return oauth.ErrContextCancelled
	default:
		// continue processing
	}

	if !utility.Exactly(req.GetClient().GetResponseTypes(), oauth.ResponseTypeCode) {
		return h.next.Authorize(req, resp, ctx)
	}

	if !h.ScopeStrategy.AcceptsAll(req.GetClient(), req.GetSession().GetGrantedScopes()) {
		return oauth.ErrBadScope
	}

	if code, err := h.CodeStrategy.NewCode(req, ctx); err != nil {
		return err
	} else if err := h.CodeStorage.Save(code, req, ctx); err != nil {
		return err
	} else {
		resp.SetCode(code)
	}

	req.HandledResponseType(oauth.ResponseTypeCode)

	if h.next != nil {
		return h.next.Authorize(req, resp, ctx)
	}
	return nil
}

