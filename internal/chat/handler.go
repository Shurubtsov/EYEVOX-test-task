package chat

import (
	"net/http"

	"github.com/dshurubtsov/internal/handlers"
	"github.com/dshurubtsov/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler() handlers.Handler {
	logger := logging.GetLogger()
	return &handler{logger: logger}
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
	w.Write([]byte("Create chat!"))
	w.WriteHeader(http.StatusCreated)
}
