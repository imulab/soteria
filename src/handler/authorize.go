package handler

import (
	"context"
	"encoding/json"
	"github.com/imulab/soteria/pkg/oauth"
	"github.com/imulab/soteria/pkg/oauth/handler"
	"github.com/imulab/soteria/pkg/oauth/request"
	"net/http"
	"time"
)

type AuthorizeHandler struct {
	ParserChain  request.OAuthAuthorizeRequestParser
	HandlerChain handler.OAuthAuthorizeHandler
}

func (h *AuthorizeHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancelFunc()

	req := request.NewOAuthAuthorizeRequest()
	resp := oauth.NewDefaultAuthorizeResponse()

	select {
	case <-ctx.Done():
		h.RenderError(rw, r, req, resp, ctx.Err())
	default:
		if err := h.ParserChain.Parse(ctx, r, req); err != nil {
			h.RenderError(rw, r, req, resp, err)
		} else if err := h.HandlerChain.Authorize(req, resp, ctx); err != nil {
			h.RenderError(rw, r, req, resp, err)
		} else {
			h.RenderResponse(rw, r, req, resp)
		}
	}
}

func (h *AuthorizeHandler) RenderError(rw http.ResponseWriter, r *http.Request, req request.OAuthAuthorizeRequest,
	resp oauth.AuthorizeResponse, err error) {
		jsonBytes, _ := json.Marshal(struct {
			Error 			string	`json:"error"`
			ErrorMessage	string	`json:"error_message"`
		}{
			Error: err.Error(),
		})
		rw.Write(jsonBytes)
		rw.WriteHeader(500)
}

func (h *AuthorizeHandler) RenderResponse(rw http.ResponseWriter, r *http.Request, req request.OAuthAuthorizeRequest,
	resp oauth.AuthorizeResponse) {
	rw.Write([]byte("hello"))
	rw.WriteHeader(200)
}