package main

import (
	"library-app/internal/db"
	"log/slog"
	"net/http"
)

func main() {
	_, err := db.InitDB()
	if err != nil {
		slog.Error("err", err)
		return
	}

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("err", err)
		return
	}

}
