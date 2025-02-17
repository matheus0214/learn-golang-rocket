package main

import (
	"log/slog"
	"net/http"
	"os"
	"shortenURLWithDB/internal/api"
	"shortenURLWithDB/internal/store"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Error to run", "errr", err)
		os.Exit(1)
	}

	slog.Info("All processing is finished")
}

func run() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	store := store.NewStore(rdb)

	handler := api.NewHandler(store)

	s := http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      handler,
	}

	slog.Info("running server", "port", 8080)

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
