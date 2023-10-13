package main

import (
	"context"
	"employee/internal/config"
	"employee/internal/server"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func startServer(e *echo.Echo) {
	go func() {
		if err := e.Start(":3000"); err != nil && err != http.ErrServerClosed {
			log.Error(err)
			e.Logger.Fatal("shutting down the server")
		}
	}()
}

func shutDownServer(e *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	log.Println("Server shutdown gracefully")
}

func main() {

	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.JSONFormatter{})
	cfg := config.NewConfig()

	sqlConn := config.GetDBInstance(*cfg)
	srv := server.NewServer(cfg, sqlConn)
	srv.ConfigureRoutes()

	startServer(srv.Echo)
	shutDownServer(srv.Echo)
}
