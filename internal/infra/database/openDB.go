package database

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/configs"
	_ "github.com/lib/pq"
)

var db *sql.DB

func OpenDB() {
	var databaseConfig = configs.LoadDatabaseConfig()
    db, err := sql.Open("postgres", databaseConfig.Dsn)
	if err != nil {
        panic(err.Error())
	}
	db.SetMaxOpenConns(databaseConfig.MaxOpenConns)
	db.SetMaxIdleConns(databaseConfig.MaxIdleConns)
	db.SetConnMaxIdleTime(databaseConfig.MaxIdleTime)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pingErr := db.PingContext(ctx)
	if pingErr != nil {
		db.Close()
        panic(pingErr.Error())
	}
	slog.Info("Database connection pool established")
}

func GetDB() *sql.DB {
    return db
}
