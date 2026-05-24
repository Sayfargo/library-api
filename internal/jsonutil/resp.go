package jsonutil

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func EncodeJSON(w http.ResponseWriter, v any, statusCode ...int) error {

	w.Header().Set("Content-Type", "application/json")

	buf, err := json.Marshal(v)
	if err != nil {
		slog.Warn("Не удалось энкодировать JSON", "err", err)
		http.Error(w, "Внутреняя ошибка сервера", http.StatusBadRequest)
		return err
	}

	if len(statusCode) > 0 {
		w.WriteHeader(statusCode[0])
	} else {
		w.WriteHeader(http.StatusOK)
	}

	_, err = w.Write(buf)

	return err
}

// EncodeSuccess - успешный ответ
func EncodeSuccess(w http.ResponseWriter, data any, status int) error {
	return EncodeJSON(w, map[string]any{
		"success": true,
		"data":    data,
	}, status)
}

// EncodeError - ответ с ошибкой
func EncodeError(w http.ResponseWriter, message string, status int) error {
	return EncodeJSON(w, map[string]any{
		"success": false,
		"error":   message,
	}, status)
}
