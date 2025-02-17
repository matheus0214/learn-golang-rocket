package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func sendJSON(w http.ResponseWriter, r apiResponse, s int) {
	data, err := json.Marshal(r)
	if err != nil {
		sendJSON(w, apiResponse{Error: "erro to parse data"}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	if _, err := w.Write(data); err != nil {
		slog.Error("Erro to send data to client", "error", err)
		return
	}
}
