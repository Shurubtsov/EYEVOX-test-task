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

// Initialize API endpoints
func (h *handler) Register(router *httprouter.Router) {
	router.POST("/chats/create", h.CreateChat) // endpoint for create chat
}

func (h *handler) CreateChat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// create entity
	chat := Chat{}

	// read request body for unmarshal json to struct of chat entity
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//h.logger.Errorf("Error with get body from request, incorrect data, err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.Unmarshal(data, &chat)

	// creating chat
	if err = h.service.CreateChat(context.TODO(), &chat); err != nil {
		//h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// respond
	w.WriteHeader(http.StatusCreated)
}
