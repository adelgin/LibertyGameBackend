package postgres

import (
	"database/sql"
	"log"
	"os"

	"libertyGame/config"
	"libertyGame/pkg/logger"

	_ "github.com/lib/pq"
)

func Parse() config.Db {
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
