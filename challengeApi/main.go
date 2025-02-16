package main

import (
	"challengeApi/api"
	"challengeApi/database"
	"log/slog"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Error", "error", err)
		return
	}

	slog.Info("All processes is finished")
}

func run() error {
	db := database.NewDatabase()

	if err := api.Serve(db); err != nil {
		return err
	}
	return nil
}
