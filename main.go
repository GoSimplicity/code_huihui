package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GoSimplicity/code_huihui/pkg/di"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func main() {
	injector := di.MustInitContainer()
	defer func(i *do.Injector) {
		err := di.ShutdownContainer(i)
		if err != nil {
			fmt.Println("Shutdown container failed", err)
		}
	}(injector)
	v := di.GetViper(injector)
	mode := v.GetString("server.mode")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	logger := di.GetLogger(injector)
	logger.Info("Application starting...")
	server := di.InitGinServer(injector)
	port := v.GetString("server.port")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      server,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		logger.Info(fmt.Sprintf("Server starting on %s", addr))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(fmt.Sprintf("Server failed to start: %v", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	logger.Info("Server exited")
}
