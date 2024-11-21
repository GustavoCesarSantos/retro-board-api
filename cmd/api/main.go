package main

import (
	"log/slog"
	"os"

	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/database"
	http "github.com/GustavoCesarSantos/retro-board-api/internal/infra/http"
)

func main() {
	db, openDBErr := database.OpenDB()
	if openDBErr != nil {
		slog.Error(openDBErr.Error())
		os.Exit(1)
	}
	defer db.Close()
	slog.Info("Database connection pool established")
	err := http.Server()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
