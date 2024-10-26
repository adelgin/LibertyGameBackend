package postgres

import (
	"database/sql"
	"log"
	"os"
	"time"

	"libertyGame/config"
	"libertyGame/pkg/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ParseDB() config.Db {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("error init logger: %v", err)
	}

	dbConfig := config.Db{
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		User:         os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		Name:         os.Getenv("DB_NAME"),
		MaxOpenConns: 25,
		MaxIdleConns: 10,
	}

	connString := dbConfig.ConnectionString()

	db, err := sql.Open("postgres", connString)
	if err != nil {
		logger.Error().Err(err).Msg("Error opening connection.")
		return dbConfig
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Error().Err(err).Msg("Connection check error.")
		return dbConfig
	}

	logger.Info().Err(err).Msg("Connection to the database established successfully.")
	return dbConfig
}

// Postgres - .
type Postgres struct {
	*sqlx.DB
}

func New(cfg *config.Db) (*Postgres, error) {
	db, err := sqlx.Open("pgx", cfg.ConnectionString())
	if err != nil {
		return nil, err
	}

	pgdb := Postgres{
		db,
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifeTime * time.Minute)

	return &pgdb, nil
}
