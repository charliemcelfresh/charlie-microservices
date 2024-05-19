package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/joho/godotenv"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var (
	logger    *slog.Logger
	oneLogger sync.Once
	db        *sqlx.DB
	oneDB     sync.Once
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func DB() *sqlx.DB {
	oneDB.Do(
		func() {
			var err error
			dbUrl := os.Getenv("DATABASE_URL")
			connStr := fmt.Sprintf("%v", dbUrl)
			db, err = sqlx.Open("postgres", connStr)
			if err != nil {
				panic(err)
			}
			db.SetConnMaxLifetime(0)
			db.SetMaxIdleConns(3)
			db.SetMaxOpenConns(3)
		},
	)
	return db
}

func GetLogger() *slog.Logger {
	oneLogger.Do(
		func() {
			logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
		},
	)
	return logger
}
