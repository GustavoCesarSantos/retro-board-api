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
	database.OpenDB()
    oauth2.SetProvider()
	err := http.Server()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
