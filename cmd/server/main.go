package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aithlete/aithlete-api/infrastructure/logger"
	"github.com/aithlete/aithlete-api/interfaces/router"
	"github.com/aithlete/aithlete-api/pkg/app"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	log := logger.New()
	deps := app.Bootstrap(log)

	e := router.New(log, deps.Handlers)

	cfg := deps.Config.Server
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	go func() {
		log.Info("Environment: %s", cfg.Env)
		if cfg.TLS.Enabled {
			log.Info("Starting server with TLS on %s", addr)
			if err := e.StartTLS(addr, cfg.TLS.CertFile, cfg.TLS.KeyFile); err != nil && err != http.ErrServerClosed {
				log.Error("Server failed to start: %v", err)
				os.Exit(1)
			}
		} else {
			log.Info("Starting server on %s", addr)
			if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
				log.Error("Server failed to start: %v", err)
				os.Exit(1)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown: %v", err)
		os.Exit(1)
	}

	log.Info("Server exited properly")
}
