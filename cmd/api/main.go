package main

import (
	"log/slog"
	"os"

	http "github.com/GustavoCesarSantos/retro-board-api/internal/infra/http"
)

func main() {
	err := http.Server()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
