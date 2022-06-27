package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/dshurubtsov/internal/chat"
	"github.com/dshurubtsov/internal/message"
	"github.com/dshurubtsov/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

func main() {

	logging.Init()
	logger := logging.GetLogger()

	router := httprouter.New()
	chatHandler := chat.NewHandler()
	messageHandler := message.NewHandler()

	logger.Info("Register handlers")
	chatHandler.Register(router)
	messageHandler.Register(router)

	start(router, &logger)
}

func start(router *httprouter.Router, logger *logging.Logger) {
	logger.Info("Start application")
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Server is listening on port :4000")
	log.Fatalln(server.Serve(listener))
}
