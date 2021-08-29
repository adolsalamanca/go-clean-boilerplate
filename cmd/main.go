package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/adolsalamanca/go-rest-boilerplate/internal/application"
	"github.com/adolsalamanca/go-rest-boilerplate/internal/infrastructure/config"
	"github.com/adolsalamanca/go-rest-boilerplate/internal/infrastructure/environment"
	_interface "github.com/adolsalamanca/go-rest-boilerplate/internal/interface"
)

func main() {
	cfg := config.LoadConfigProvider()
	err := environment.Verify(cfg)
	if err != nil {
		log.Fatalf("could not initialize app: %v", err)
	}
	svc := _interface.NewService(cfg)
	Run(svc)
}

func Run(svc *_interface.Service) {
	_, cancelFunc := context.WithCancel(context.Background())
	server := application.NewServer(svc)

	go func() {
		fmt.Printf("Starting server... \n")
		if err := http.ListenAndServe(":8080", server.Routes()); err != nil {
			log.Fatalf("could not initialize server: %v", err)
		}
	}()

	arrangeGracefullyShutdown(cancelFunc)
}

func arrangeGracefullyShutdown(cancelFunc context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-signalChan

	log.Printf("Shutting app...")
	cancelFunc()
	os.Exit(1)
}
