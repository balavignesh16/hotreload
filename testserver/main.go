package main

import (
	"hotreload/testserver/api"
	"log/slog"
	"net/http"
	"os"
)

type Response struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Message string `json:"message"`
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	http.HandleFunc("/api/health", api.HealthCheck)
	http.HandleFunc("/api/users", api.GetUsers)

	slog.Info("API server is up")
	slog.Info("url http://localhost:8080/api/health")
	slog.Info("url http://localhost:8080/api/users")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("server crashed", "error", err)
		os.Exit(1)
	}
}
