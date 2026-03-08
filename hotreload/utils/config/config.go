package config

import (
	"errors"
	"flag"
	"log/slog"
	"os"
)

type Config struct {
	RootPath     string
	BuildCommand string
	ExecCommand  string
}

func ParseFlags() (*Config, error) {
	cfg := &Config{}

	flag.StringVar(&cfg.RootPath, "root", "", "Directory to watch for file changes")
	flag.StringVar(&cfg.BuildCommand, "build", "", "Command used to build the project")
	flag.StringVar(&cfg.ExecCommand, "exec", "", "Command used to run the built server")

	flag.Parse()

	if cfg.RootPath == "" {
		err := errors.New("--root flag is required")
		slog.Error("missing required flag", "flag", "root", "error", err)
		return nil, err
	}

	if cfg.BuildCommand == "" {
		err := errors.New("--build flag is required")
		slog.Error("missing required flag", "flag", "build", "error", err)
		return nil, err
	}

	if cfg.ExecCommand == "" {
		err := errors.New("--exec flag is required")
		slog.Error("missing required flag", "flag", "exec", "error", err)
		return nil, err
	}

	info, err := os.Stat(cfg.RootPath)
	if err != nil {
		slog.Error("failed to stat root path", "path", cfg.RootPath, "error", err)
		return nil, err
	}

	if !info.IsDir() {
		err := errors.New("provided --root path is not a directory")
		slog.Error("invalid root path", "path", cfg.RootPath, "error", err)
		return nil, err
	}

	slog.Info("configuration loaded",
		"root", cfg.RootPath,
		"build", cfg.BuildCommand,
		"exec", cfg.ExecCommand,
	)

	return cfg, nil
}
