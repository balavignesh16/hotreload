package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Message string `json:"message"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	slog.Info("Request received", "method", r.Method)

	resp := Response{
		Status:  "Good",
		Version: "v1",
		Message: "hotreload",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
