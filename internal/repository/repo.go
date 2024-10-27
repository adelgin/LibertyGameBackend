package repository

import (
	"context"
	//"database/sql"
	"fmt"
	postgres "libertyGame/internal"
	"time"

	_ "github.com/goccy/go-json"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	GetUserByID(ctx context.Context, id int64) (*User, error)
	AddUser(ctx context.Context, user *User) error
	CountOfAllUsers(ctx context.Context) (int64, error)
	GetRefsOfUserFromID(ctx context.Context, id int64) ([]User, error)
	CountRefsOfUserFromID(ctx context.Context, id int64) (int64, error)
	GetTopOfRefs(ctx context.Context, count int64) ([]Top_User, error)
	GetMonthStatistics(ctx context.Context) ([]MonthStatistics, error)
	CreateTable(ctx context.Context) error
}

type User struct {
	UserID    int64     `db:"id"             json:"id"`
	UserName  string    `db:"username"       json:"username"`
	InviterID *int64    `db:"inviter_id"     json:"inviter_id,omitempty"`
	CreatedAt time.Time `db:"created_at"     json:"date_of_invite"`
}

type Top_User struct {
	UserID       int64     `db:"id"             json:"id"`
	UserName     string    `db:"username"       json:"username"`
	InviterCount *int64    `db:"inviter_count"  json:"inviter_count"`
	CreatedAt    time.Time `db:"created_at"     json:"date_of_invite"`
}

type MonthStatistics struct {
	Months    string `db:"month"             json:"month"`
	UserCount int64  `db:"user_count"             json:"user_count"`
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) GetUserByID(ctx context.Context, id int64) (*User, error) {
	var user User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r repository) AddUser(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (id, username, inviter_id, created_at)
	VALUES (:id, :username, :inviter_id, :created_at)
	RETURNING id;
	`

	_, err := r.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}

	return nil
}

func (r repository) CountOfAllUsers(ctx context.Context) (int64, error) {
	var counter int64
	err := r.db.GetContext(ctx, &counter, "SELECT COUNT(*) FROM users")
	if err != nil {
		return 0, err
	}
	return counter, nil
}

func (r repository) GetRefsOfUserFromID(ctx context.Context, id int64) ([]User, error) {
	var users []User
	err := r.db.SelectContext(ctx, &users, "SELECT * FROM users WHERE inviter_id = $1", id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r repository) CountRefsOfUserFromID(ctx context.Context, id int64) (int64, error) {
	var counter int64
	err := r.db.GetContext(ctx, &counter, "SELECT COUNT(*) FROM users WHERE inviter_id = $1;", id)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

func (r repository) GetTopOfRefs(ctx context.Context, count int64) ([]Top_User, error) {
	var users []Top_User
	rows, err := r.db.QueryContext(ctx, "WITH InvitedCounts AS (SELECT u.id, u.username, COUNT(inv.id) AS invited_count, u.created_at FROM users u LEFT JOIN users inv ON u.id = inv.inviter_id GROUP BY u.id, u.username, u.inviter_id, u.created_at) SELECT * FROM InvitedCounts ORDER BY invited_count DESC LIMIT $1;", count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user Top_User

		if err := rows.Scan(&user.UserID, &user.UserName, &user.InviterCount, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r repository) GetMonthStatistics(ctx context.Context) ([]MonthStatistics, error) {
	var stats []MonthStatistics
	query := "SELECT TO_CHAR(created_at, 'MM.YYYY') AS month, COUNT(DISTINCT id) AS user_count FROM users GROUP BY month ORDER BY month"
	err := r.db.SelectContext(ctx, &stats, query)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func (r repository) CreateTable(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS users (
	id BIGINT PRIMARY KEY,
	username VARCHAR(255) UNIQUE NOT NULL,
	inviter_id INT,
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (inviter_id) REFERENCES users(id)
	)`
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы users: %w", err)
	}
	return nil
}
