package main

import (
	"hotreload/utils/config"
	"log/slog"
	"os"
)

type Config struct {
	RootPath     string
	BuildCommand string
	ExecCommand  string
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cfg, err := config.ParseFlags()
	if err != nil {
		slog.Error("Failed to parse configuration", "error", err)
		os.Exit(1)
	}

}
