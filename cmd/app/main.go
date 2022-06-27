package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/dshurubtsov/internal/chat"
	"github.com/dshurubtsov/internal/message"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	chatHandler := chat.NewHandler()
	messageHandler := message.NewHandler()

	log.Println("Register handlers")
	chatHandler.Register(router)
	messageHandler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("Start application")
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("server is listening on port :4000")
	log.Fatalln(server.Serve(listener))
}
