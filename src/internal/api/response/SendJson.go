package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func SendJSON(w http.ResponseWriter, resp Data, status int) {
	data, err := json.Marshal(resp)

	if err != nil {
		slog.Error("failed to marshal json data", "error", err, "response", resp)
		SendJSON(w, Data{Error: "something went wrong"}, http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}
}

type Data struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}
