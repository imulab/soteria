package client

import (
	"github.com/gorilla/mux"
	"github.com/imulab/soteria/pkg/oauth/client"
	"github.com/imulab/soteria/src/handler"
	"github.com/urfave/negroni"
	"net/http"
)

type clientApi struct {
	Handler 	*handler.ClientHandler
}

func (api *clientApi) setup() error {
	api.Handler = &handler.ClientHandler{
		Repo: &client.NotFoundClientLookup{},
	}
	return nil
}

func (api *clientApi) startWebServer() error {
	r := mux.NewRouter()
	r.HandleFunc("/client/{client_id}", api.Handler.HandleGet)

	n := negroni.Classic()
	n.UseHandler(r)

	return http.ListenAndServe(":8080", n)
}