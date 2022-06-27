package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/dshurubtsov/internal/chat"
	"github.com/dshurubtsov/internal/config"
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

	cfg := config.GetConfig()

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()

	logger.Info("Start application")

	logger.Infof("Listen type is on %s", cfg.Listen.Type)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	if err != nil {
		logger.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infof("Server is listening on port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	log.Fatalln(server.Serve(listener))
}
