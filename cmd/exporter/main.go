package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"exporter/internal/infrastructure/finnhub"
	"exporter/internal/interfaces/client"
	"exporter/internal/interfaces/handlers"
	"exporter/internal/usecases/collector"

	"github.com/D3vR4pt0rs/logger"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	polygonApiKey = "POLYGON_API_KEY"
	finnhubApiKey = "FINNHUB_API_KEY"
)

func main() {
	token := os.Getenv(finnhubApiKey)

	finnhubClient := finnhub.New(token)
	apiClient := client.New(finnhubClient)
	application := collector.New(apiClient)

	router := mux.NewRouter()
	handlers.Make(router, application)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},   // All origins
		AllowedMethods: []string{"GET"}, // Allowing only get, just an example
	})

	srv := &http.Server{
		Addr:    ":1338",
		Handler: c.Handler(router),
	}

	go func() {
		listener := make(chan os.Signal, 1)
		signal.Notify(listener, os.Interrupt, syscall.SIGTERM)
		fmt.Println("Received a shutdown signal:", <-listener)

		if err := srv.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
			logger.Error.Println("Failed to gracefully shutdown ", err)
		}
	}()

	logger.Info.Println("[*]  Listening...")
	if err := srv.ListenAndServe(); err != nil {
		logger.Error.Println("Failed to listen and serve ", err)
	}

	logger.Critical.Println("Server shutdown")
}
