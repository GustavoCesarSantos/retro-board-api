package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/configs"
	_ "github.com/lib/pq"
)



func OpenDB() (*sql.DB, error) {
	var databaseConfig = configs.LoadDatabaseConfig()
	db, err := sql.Open("postgres", databaseConfig.Dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(databaseConfig.MaxOpenConns)
	db.SetMaxIdleConns(databaseConfig.MaxIdleConns)
	db.SetConnMaxIdleTime(databaseConfig.MaxIdleTime)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pingErr := db.PingContext(ctx)
	if pingErr != nil {
		db.Close()
		return nil, pingErr
	}
	return db, nil
}