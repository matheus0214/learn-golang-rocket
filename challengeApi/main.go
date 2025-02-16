package main

import (
	"challengeApi/api"
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
	if err := api.Serve(); err != nil {
		return err
	}
	return nil
}
