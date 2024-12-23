package main

import (
	"context"
	"log"
	"microservice/infra"
	"microservice/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	r := routes.NewRoutes(*ctx)

	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	go func() {
		// Start the server
		log.Println("Server running on port " + os.Getenv("PORT"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// Create a timeout context for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// Catching context timeout
	select {
	case <-shutdownCtx.Done():
		log.Println("Timeout of 5 seconds.")
	}

	log.Println("Server exiting")
}
