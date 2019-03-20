package handler

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/imulab/soteria/pkg/oauth/client"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type ClientHandler struct {
	Repo 	client.Repository
}

func (h *ClientHandler) HandleGet(rw http.ResponseWriter, r *http.Request) {
	clientId := mux.Vars(r)["client_id"]

	logrus.WithFields(logrus.Fields{
		"client_id": clientId,
	}).Debugln("received request.")

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancelFunc()

	if c, err := h.Repo.Find(ctx, clientId); err != nil {
		h.RenderError(rw, r, err)
	} else {
		h.RenderClient(rw, r, c)
	}
}

func (h *ClientHandler) RenderError(rw http.ResponseWriter, r *http.Request, err error) {
	rw.Header().Set("Content-Type", "application/json")
}

func (h *ClientHandler) RenderClient(rw http.ResponseWriter, r *http.Request, c client.Client) {
	rw.Write([]byte("hello"))
}