package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/pkg/errors"
)

type User struct {
	UserID    int64     `db:"id"             json:"-"`
	UserName  *string   `db:"username"       json:"-"`
	InviterID int64     `db:"inviter_id"     json:"-"`
	CreatedAt time.Time `db:"created_at"     json:"-"`
}

type DBTX interface {
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	DriverName() string
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	Preparex(query string) (*sqlx.Stmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type repository struct {
	db  DBTX
	con *sqlx.DB
}

// NewRepository ...
// func NewRepository(db DBTX, con *sqlx.DB) Repository {
// 	return &repository{
// 		db:  db,
// 		con: con,
// 	}
// }

func (r *repository) GetUserByID(ctx context.Context, id int64) (User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	res := User{}

	err := r.db.GetContext(ctx, &res, query, id)

	if err != nil {
		fmt.Println(err)
		return User{}, nil //errors.Wrap(err, "cannot get user")
	}

	return res, nil
}
