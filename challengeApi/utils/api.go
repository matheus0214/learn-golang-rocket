package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func JsonResponse(w http.ResponseWriter, r Response, s int) {
	data, err := json.Marshal(r)
	if err != nil {
		JsonResponse(w, Response{Error: "Erro to parse response"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(s)

	if _, err := w.Write(data); err != nil {
		slog.Error("Erro to send data to client", "error", err)
		return
	}
}
