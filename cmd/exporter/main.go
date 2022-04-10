package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"exporter/internal/infrastructure/polygon"
	"exporter/internal/interfaces/client"
	"exporter/internal/interfaces/handlers"
	"exporter/internal/usecases/collector"

	"github.com/gorilla/mux"
)

const (
	polygonApiKey = "POLYGON_API_KEY"
)

func main() {
	token := os.Getenv(polygonApiKey)

	polygonClient := polygon.New(token)
	apiClient := client.New(polygonClient)
	application := collector.New(apiClient)

	router := mux.NewRouter()
	handlers.Make(router, application)
	srv := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	go func() {
		listener := make(chan os.Signal, 1)
		signal.Notify(listener, os.Interrupt, syscall.SIGTERM)
		fmt.Println("Received a shutdown signal:", <-listener)

		if err := srv.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
			fmt.Println("Failed to gracefully shutdown ", err)
		}
	}()

	fmt.Println("[*]  Listening...")
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Failed to listen and serve ", err)
	}

	fmt.Println("Server shutdown")
}
