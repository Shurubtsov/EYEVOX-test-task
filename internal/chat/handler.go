package chat

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dshurubtsov/internal/handlers"
	"github.com/dshurubtsov/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	logger  *logging.Logger
	service ChatService
}

func NewHandler(service ChatService) handlers.Handler {
	logger := logging.GetLogger()
	return &handler{logger: logger, service: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/chats", h.GetChats)           // *. получение листа чатов
	router.POST("/chats/create", h.CreateChat) // 1. создание чата
}

func (h *handler) GetChats(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is list of "))
	w.WriteHeader(http.StatusOK)
}

func (h *handler) CreateChat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	chat := Chat{}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Errorf("Error with get body from request, incorrect data, err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.Unmarshal(data, &chat)

	if err = h.service.CreateChat(context.TODO(), &chat); err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
