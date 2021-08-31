package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/adolsalamanca/go-clean-boilerplate/internal/application"
	_interface "github.com/adolsalamanca/go-clean-boilerplate/internal/interface"
	"github.com/adolsalamanca/go-clean-boilerplate/pkg/config"
	log "github.com/adolsalamanca/go-clean-boilerplate/pkg/logger"
	"github.com/adolsalamanca/go-clean-boilerplate/pkg/metrics"
)

func main() {
	logger := log.NewLogger(log.DebugLevel)

	cfg := config.LoadConfigProvider()
	err := _interface.Verify(cfg, logger)
	if err != nil {
		logger.Error("could not initialize app: %v", log.NewFieldString("error", err.Error()))
	}

	statsdAddress := fmt.Sprintf("%s:%d", cfg.GetString("STATSD_HOST"), cfg.GetInt("STATSD_PORT"))
	collector, err := metrics.NewMetricsCollector(statsdAddress, "go_rest_boilerplate")

	Run(cfg, logger, collector)
}

func Run(cfg config.Provider, logger _interface.Logger, collector _interface.MetricsCollector) {

	svc, err := _interface.NewService(cfg, logger, collector)
	if err != nil {
		logger.Error("could not create service", log.NewFieldString("error", err.Error()))
	}
	server := application.NewServer(svc, logger, collector)

	go func() {
		logger.Info("Starting server")
		if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.GetString("SERVER_PORT")), server.Routes()); err != nil {
			logger.Error("could not initialize app: %v", log.NewFieldString("error", err.Error()))
		}
	}()

	_, cancelFunc := context.WithCancel(context.Background())
	arrangeGracefullyShutdown(cancelFunc, logger)
}

func arrangeGracefullyShutdown(cancelFunc context.CancelFunc, logger _interface.Logger) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-signalChan

	logger.Info("shutting down app")
	cancelFunc()
	os.Exit(1)
}
