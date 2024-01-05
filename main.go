package main

import (
	"context"
	"errors"
	"go-server/internal/config"
	"go-server/internal/storage/postgres"
	"go-server/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func main() {
	config.MustLoad()
	config.InitLogger()

	postgres.ConnectDb()
	r := routes.SetupRouter()

	start(r)
}

func start(r *chi.Mux) {
	var err error
	port := viper.GetString("PORT")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	config.Logger.Info("Starting server on port " + port)

	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
