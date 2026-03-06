package main

import (
	"encoding/json"
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

	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Request received", "method", r.Method)
		resp := Response{
			Status:  "Good",
			Version: "v1",
			Message: "hotreload",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

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
