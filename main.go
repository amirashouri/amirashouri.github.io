package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"portfolio/handlers"
	"portfolio/middleware"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	middlewareStack := middleware.CreateStack(middleware.TextHTMLMiddleware, middleware.CSPMiddleware)
	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	router.HandleFunc("/", handlers.NewHomeHandler().ServeHTTP)

	killSig := make(chan os.Signal, 1)

	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)
	listenAddr := os.Getenv("LISTEN_ADDR")

	srv := &http.Server{
		Addr:    listenAddr,
		Handler: middlewareStack(router),
	}

	go func() {

		err := srv.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			slog.Info("Server shutdown complete")
		} else if err != nil {
			slog.Error("Server error", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	slog.Info("Server started...")
	<-killSig

	slog.Info("Shutting down server")

	// Create a context with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("Server shutdown complete")
}
