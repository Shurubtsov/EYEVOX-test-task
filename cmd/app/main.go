package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/dshurubtsov/internal/chat"
	chatdb "github.com/dshurubtsov/internal/chat/db"
	"github.com/dshurubtsov/internal/config"
	"github.com/dshurubtsov/internal/message"
	msgdb "github.com/dshurubtsov/internal/message/db"
	"github.com/dshurubtsov/pkg/client/postgresql"
	"github.com/dshurubtsov/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

func main() {

	// parent context of app
	ctx := context.Background()

	// logger initialization
	logging.Init()
	logger := logging.GetLogger()

	// get configuration for app
	cfg := config.GetConfig()
	logger.Info("Config: ", cfg)

	// init router for handlers
	router := httprouter.New()

	postgreClient, err := postgresql.NewClient(ctx, 5, cfg.Storage, logger)
	if err != nil {
		logger.Fatalf("Can't create client of postgresql, err: %v", err)
	}

	// chat entity
	chatRepository := chatdb.NewRepository(postgreClient, logger)
	chatService := chat.NewService(chatRepository, logger)
	chatHandler := chat.NewHandler(chatService, ctx)

	// message entity
	msgRepository := msgdb.NewRepository(postgreClient, logger)
	msgService := message.NewService(msgRepository, logger)
	messageHandler := message.NewHandler(msgService, ctx)

	logger.Info("Register handlers")

	// registration our handlers for entities from router
	chatHandler.Register(router)
	messageHandler.Register(router)

	// Start app
	start(router, cfg)
}

// Entry point for start app
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

	logger.Infof("Start server listening on port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	log.Fatalln(server.Serve(listener))
}
