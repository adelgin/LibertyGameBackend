package repository

import (
	"context"
	"fmt"
	postgres "libertyGame/internal"
	"time"

	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	UserID    int64     `db:"id"             json:"-"`
	UserName  string    `db:"username"       json:"-"`
	InviterID int64     `db:"inviter_id"     json:"-"`
	CreatedAt time.Time `db:"created_at"     json:"-"`
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db postgres.Postgres) repository {
	return repository{
		db: &db,
	}
}

func (r repository) GetUserByID(ctx context.Context, id int64) (*User, error) { // tested
	var user User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r repository) AddUser(ctx context.Context, user *User) error { //tested
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

func (r repository) CountOfAllUsers(ctx context.Context) (int64, error) { //tested
	var counter int64
	err := r.db.GetContext(ctx, &counter, "SELECT COUNT(*) FROM users")
	if err != nil {
		return 0, err
	}
	return counter, nil
}

func (r repository) GetRefsOfUserFromID(ctx context.Context, id int64) ([]User, error) { //tested
	var users []User
	err := r.db.SelectContext(ctx, &users, "SELECT * FROM users WHERE inviter_id = $1", id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r repository) CountRefsOfUserFromID(ctx context.Context, id int64) (int64, error) { //tested
	var counter int64
	err := r.db.GetContext(ctx, &counter, "SELECT COUNT(*) FROM users WHERE inviter_id = $1;", id)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

func (r repository) GetTopOfRefs(ctx context.Context, count int64) ([]User, error) { //tested
	var users []User
	rows, err := r.db.QueryContext(ctx, "SELECT u.id, u.username, COUNT(r.id) AS invites_count FROM users u LEFT JOIN users r ON u.id = r.inviter_id GROUP BY u.id, u.username ORDER BY invites_count DESC LIMIT $1", count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		var invitesCount int64

		if err := rows.Scan(&user.UserID, &user.UserName, &invitesCount); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
