package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"

	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/database"
	http "github.com/GustavoCesarSantos/retro-board-api/internal/infra/http"
	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/oauth2"
)

func main() {
	loadEnvErr := godotenv.Load()
	if loadEnvErr != nil {
		slog.Error("failed to load .env file", "error", loadEnvErr)
        os.Exit(1)
	}
	DB, DBErr := database.OpenDB()
	if DBErr != nil {
		slog.Error(DBErr.Error())
		os.Exit(1)
	}
	defer DB.Close()
    oauth2.SetProvider()
	serverErr := http.Server(DB)
	if serverErr != nil {
		slog.Error(serverErr.Error())
		os.Exit(1)
	}
}
