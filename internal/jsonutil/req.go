package jsonutil

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func DecodeJSON(w http.ResponseWriter, r *http.Request, v any) error {

	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024)

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		slog.Warn("ошибка декодирования JSON", "err", err)
		EncodeError(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}
