package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sjc5/go-api-template/internal/platform"
	"github.com/sjc5/go-api-template/internal/router"
)

func main() {
	// Create a new router and setup the API
	r := router.Init()

	// Initialize the server
	server := &http.Server{Addr: "0.0.0.0:" + platform.GetEnv().Port, Handler: r}

	// Setup the server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown the server with a grace period
		shutdownCtx, cancelCtx := context.WithTimeout(serverCtx, time.Duration(platform.GetEnv().GracefulTimeoutSeconds)*time.Second)
		defer cancelCtx()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out. forcing exit.")
			}
		}()

		// Trigger the server shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Start the server
	fmt.Printf("starting server on: http://localhost:%s\n", platform.GetEnv().Port)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for the server to stop
	<-serverCtx.Done()
}
