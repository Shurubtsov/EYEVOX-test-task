package chat

import (
	"net/http"

	"github.com/dshurubtsov/internal/handlers"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/chats", h.GetChats)           // *. получение листа чатов
	router.POST("/chats/create", h.CreateChat) // 1. создание чата
}

func (h *handler) GetChats(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is list of "))
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) CreateChat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Create chat!"))
}
