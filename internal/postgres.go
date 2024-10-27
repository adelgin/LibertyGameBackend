package postgres

import (
	"time"

	"libertyGame/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

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
