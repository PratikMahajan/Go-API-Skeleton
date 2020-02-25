package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/PratikMahajan/Go-Twitter-Downloader-Bot/config"
	"github.com/PratikMahajan/Go-Twitter-Downloader-Bot/handlers"
	"github.com/PratikMahajan/Go-Twitter-Downloader-Bot/metrics"

	"github.com/Noah-Huppert/golog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/gorilla/mux"


)

func main(){

	// {{{1 Context
	ctx, ctxCancel := context.WithCancel(context.Background())

	// signals holds signals received by process
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals

		ctxCancel()
	}()


	// {{{1 Logger
	logger := golog.NewStdLogger("GoAppTwitter")
	logger.Debug("starting")


	// {{{1 Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("failed to load configuration: %s", err.Error())
	}

	logger.Debugf("loaded configuration:")


	// {{{1 Setup Prometheus metrics
	metricsInstance := metrics.NewMetrics()


	// {{{1 Setup shutdown wait group
	// shutdownWaitGroup is used to ensure that all components have gracefuly shut down before the process exists
	var shutdownWaitGroup sync.WaitGroup

	// {{{1 Prometheus metrics server
	metricsRouter := mux.NewRouter()
	metricsRouter.Handle("/metrics", promhttp.Handler())

	metricsServer := http.Server{
		Addr:    cfg.MetricsAddr,
		Handler: metricsRouter,
	}

	logger.Debug("starting metrics server")

	shutdownWaitGroup.Add(1)
	go func() {
		defer shutdownWaitGroup.Done()

		if err := metricsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("failed to serve metrics: %s", err.Error())
		}

		logger.Debug("stopped metrics server")
	}()

	shutdownWaitGroup.Add(1)
	go func() {
		defer shutdownWaitGroup.Done()

		<-ctx.Done()

		if err := metricsServer.Shutdown(context.Background()); err != nil {
			logger.Fatalf("failed to shutdown metrics server: %s",
				err.Error())
		}
	}()

	logger.Infof("started metrics server on %s", cfg.MetricsAddr)


	// {{{1 API Router
	baseHandler := handlers.BaseHandler{
		Ctx:     ctx,
		Logger:  logger.GetChild("handlers"),
		Cfg:     cfg,
		Metrics: metricsInstance,
	}


	apiRouter := mux.NewRouter()

	apiRouter.Handle("/health", handlers.HealthHandler{
		baseHandler.GetChild("health"),
	}).Methods("GET")





	// {{{1 Start API server
	logger.Debug("starting API server")

	apiServer := http.Server{
		Addr: cfg.APIAddr,
		Handler: handlers.PanicHandler{
			BaseHandler: baseHandler,
			Handler: handlers.MetricsHandler{
				BaseHandler: baseHandler,
				Handler: handlers.ReqLoggerHandler{
					BaseHandler: baseHandler,
					Handler:     apiRouter,
				},
			},
		},
	}

	shutdownWaitGroup.Add(1)
	go func() {
		defer shutdownWaitGroup.Done()

		if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("failed to serve API: %s", err.Error())
		}

		logger.Debug("stopped API server")
	}()

	shutdownWaitGroup.Add(1)
	go func() {
		defer shutdownWaitGroup.Done()

		<-ctx.Done()

		if err := apiServer.Shutdown(context.Background()); err != nil {
			logger.Fatalf("failed to shutdown API server: %s",
				err.Error())
		}
	}()

	logger.Infof("started API server on %s", cfg.APIAddr)

	// {{{1 Wait for all components to shut down
	shutdownWaitGroup.Wait()

	logger.Info("done")
}