package authorize

import (
	"github.com/gorilla/mux"
	"github.com/imulab/soteria/pkg/oauth/client"
	handler2 "github.com/imulab/soteria/pkg/oauth/handler"
	"github.com/imulab/soteria/pkg/oauth/request"
	"github.com/imulab/soteria/pkg/oauth/token"
	"github.com/imulab/soteria/src/handler"
	"github.com/pkg/errors"
	"github.com/urfave/negroni"
	"net/http"
)

type authorizeApi struct {
	Handler *handler.AuthorizeHandler
}

func (api *authorizeApi) setup() error {
	var err error

	// parsers
	rootParser := &request.OAuthAuthorizeRequestQueryParser{
		ClientLookup: &client.NotFoundClientLookup{},
		ClientLookupTimeoutSeconds: 10,
	}
	rootParser.WithNext(&request.OAuthAuthorizeRequestSessionParser{})

	// handlers
	codeStrategy, err := token.NewHmacSha256AuthorizeCodeStrategy(16, []byte("2530120357574159ad772a4daf2ef7ea"))
	if err != nil {
		return errors.WithStack(err)
	}
	rootHandler := &handler2.AuthorizeCodeHandler{
		ScopeStrategy: &client.EqualityScopeStrategy{IgnoreCase: false},
		CodeStorage: token.NewNoOpAuthorizeCodeRepository(),
		CodeStrategy: codeStrategy,
	}

	// bootstrap
	api.Handler = &handler.AuthorizeHandler{
		ParserChain: rootParser,
		HandlerChain: rootHandler,
	}

	return nil
}

func (api *authorizeApi) startWebServer() error {
	r := mux.NewRouter()
	r.HandleFunc("/oauth/authorize", api.Handler.Handle)

	n := negroni.Classic()
	n.UseHandler(r)

	return http.ListenAndServe(":8080", n)
}