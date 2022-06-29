package message

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dshurubtsov/internal/handlers"
	"github.com/dshurubtsov/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	logger     *logging.Logger
	serviceMsg MessageService
}

func NewHandler(service MessageService) handlers.Handler {
	logger := logging.GetLogger()
	return &handler{logger: logger, serviceMsg: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/messages/:chatName/:page", h.GetListMessages) // 3. получение списка ID сообщений
	router.GET("/message/:id", h.GetMessageByID)               // 4. получение сообщения по его ID
	router.POST("/message/create/:chatName", h.CreateMessage)  // 2. добавление сообщений в чат
}

func (h *handler) GetListMessages(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	chatName := params.ByName("chatName")
	page, err := strconv.Atoi(params.ByName("page"))
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if page < 1 {
		h.logger.Error("Page is less zero")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// pagination
	limit := 5
	offset := limit * (page - 1)

	listID, err := h.serviceMsg.FindListID(context.TODO(), chatName, limit, offset)
	if err != nil {
		h.logger.Info("Can't get list of id")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(listID)
	if err != nil {
		h.logger.Error("Can't marshal list id, error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetMessageByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	message, err := h.serviceMsg.FindMessageByID(context.TODO(), id, &Message{})
	if err != nil {
		h.logger.Error("Error found message")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(&message)
	if err != nil {
		h.logger.Error("Error marshal response to json, error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) CreateMessage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	chatName := params.ByName("chatName")
	msg := Message{}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Errorf("Error with get body from request, incorrect data, err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.Unmarshal(data, &msg)

	if msg.CreatorNickname == "" || msg.TextMessage == "" {
		h.logger.Error("Empty body request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.serviceMsg.CreateMessage(context.TODO(), &msg, chatName); err != nil {
		h.logger.Errorf("Error with create new message, err: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
