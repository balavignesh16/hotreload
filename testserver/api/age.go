// http://localhost:8080/api/age?dob=12_03_1999
package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type AgeResponse struct {
	DOB           string `json:"dob"`
	CalculatedAge int    `json:"calculated_age"`
}

func CalculateAge(w http.ResponseWriter, r *http.Request) {
	dobParam := r.URL.Query().Get("dob")
	if dobParam == "" {
		slog.Error("missing dob query")
		http.Error(w, `{"error": "missing dob query"}`, http.StatusBadRequest)
		return
	}

	dob, err := time.Parse("02_01_2006", dobParam)
	if err != nil {
		slog.Error("invalid format given", "input", dobParam, "error", err)
		http.Error(w, `{"error": "invalid format given Use dd__mm__yyyy"}`, http.StatusBadRequest)
		return
	}

	now := time.Now()
	age := now.Year() - dob.Year()

	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}

	resp := AgeResponse{
		DOB:           dobParam,
		CalculatedAge: age,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
