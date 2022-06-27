package message

import (
	"fmt"
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
	router.GET("/messages", h.GetListMessages)      // 3. получение списка ID сообщений
	router.GET("/message/:id", h.GetMessageByID)    // 4. получение сообщения по его ID
	router.POST("/message/create", h.CreateMessage) // 2. добавление сообщений в чат
}

func (h *handler) GetListMessages(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("List of id messages"))
	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetMessageByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	w.Write([]byte(fmt.Sprintf("Message with id: %s", param)))
	w.WriteHeader(http.StatusOK)
}

func (h *handler) CreateMessage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Create message."))
	w.WriteHeader(http.StatusCreated)
}
