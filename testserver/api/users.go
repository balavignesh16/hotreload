package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	slog.Info("Users list requested", "method", r.Method, "client_ip", r.RemoteAddr)

	users := []User{
		{ID: 1, Name: "Bala", Role: "Admin"},
		{ID: 2, Name: "vignesh", Role: "Software developer"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
