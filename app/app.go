package app

import (
	"context"
	"fmt"
	postgres "libertyGame/internal"
	"libertyGame/internal/repository"
	"libertyGame/pkg/logger"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"

	_ "github.com/lib/pq"
)

type User struct {
	UserID    int64     `db:"id"             json:"-"`
	UserName  string    `db:"username"       json:"-"`
	InviterID int64     `db:"inviter_id"     json:"-"`
	CreatedAt time.Time `db:"created_at"     json:"-"`
}

func Run() {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("error init logger: %v", err)
	}
	cfg := postgres.Parse()

	database, err := postgres.New(&cfg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// r := repository{
	// 	db: database,
	// }
	r := repository.NewRepository(*database)

	user, err := r.GetUserByID(context.Background(), 1226)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting user by ID")
		return
	}

	fmt.Println("Имя пользователя:", user.UserName)
	fmt.Println("ID пользователя:", user.UserID)

	// nuser := repository.User{
	// 	UserID:    1999,
	// 	UserName:  "John Hauzer",
	// 	InviterID: 1002,
	// 	CreatedAt: time.Now(),
	// }

	// err = r.AddUser(context.Background(), &nuser)
	// if err != nil {
	// 	logger.Error().Err(err).Msg("Error getting user by ID")
	// 	return
	// }

	// users, err := r.GetUserByID(context.Background(), 1234)
	// if err != nil {
	// 	logger.Error().Err(err).Msg("Error getting user by ID")
	// 	return
	// }

	// fmt.Println("Имя пользователя:", users.UserName)
	// fmt.Println("ID пользователя:", users.UserID)

	var usersarray []repository.User
	var usersarray2 []repository.User

	usersarray, err = r.GetRefsOfUserFromID(context.Background(), 1002)

	// fmt.Println(usersarray[0].UserID)
	for i := range usersarray {
		fmt.Printf("User %d %s %d %s\n", usersarray[i].UserID, usersarray[i].UserName, usersarray[i].InviterID, usersarray[i].CreatedAt)
	}

	fmt.Println(r.CountOfAllUsers(context.Background()))
	fmt.Println(r.CountRefsOfUserFromID(context.Background(), 1002))

	usersarray2, err = r.GetTopOfRefs(context.Background(), 100)
	//fmt.Println(usersarray2[0].UserID)
	for i := range usersarray2 {
		fmt.Printf("User %d %s %d %s\n", usersarray2[i].UserID, usersarray2[i].UserName, usersarray2[i].InviterID, usersarray2[i].CreatedAt)
	}

	fmt.Println(cfg.Host)
	logger.Info().Msg("Everything is fine)")
}
