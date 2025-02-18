package main

import (
	"log/slog"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config", "./internal/store/pgstore/migrations/tern.conf",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("command execution failed", "error", err)
		slog.Error("output", "output", string(output))
		panic(err)
	}

	slog.Info("command execued successfully", "output", string(output))
}
