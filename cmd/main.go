package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/adolsalamanca/go-clean-boilerplate/internal/application"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/infrastructure/config"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/infrastructure/environment"
	_interface "github.com/adolsalamanca/go-clean-boilerplate/internal/interface"
)

func main() {
	cfg := config.LoadConfigProvider()
	err := environment.Verify(cfg)
	if err != nil {
		log.Fatalf("could not initialize app: %v", err)
	}

	Run(cfg)
}

func Run(cfg config.Provider) {
	logger := _interface.NewLogger()
	service := _interface.NewService(cfg, logger)
	server := application.NewServer(service, logger)

	go func() {
		fmt.Printf("Starting server... \n")
		if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.GetString("SERVER_PORT")), server.Routes()); err != nil {
			log.Fatalf("could not initialize server: %v", err)
		}
	}()

	_, cancelFunc := context.WithCancel(context.Background())
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
