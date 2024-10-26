package app

import (
	// "fmt"
	"fmt"
	"libertyGame/config"
	postgres "libertyGame/internal"
	"libertyGame/internal/handler"
	"libertyGame/internal/repository"
	"libertyGame/internal/route"
	"libertyGame/internal/service"
	"libertyGame/internal/utils"
	chi_router "libertyGame/pkg/chi"
	"libertyGame/pkg/httpserver"
	"libertyGame/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	//_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func Run() {
	// logger, err := logger.NewLogger()
	// if err != nil {
	// 	log.Fatalf("error init logger: %v", err)
	// }
	// cfg := postgres.Parse()

	// database, err := postgres.New(&cfg)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// // r := repository{
	// // 	db: database,
	// // }
	// r := repository.NewRepository(*database)

	// user, err := r.GetUserByID(context.Background(), 1226)
	// if err != nil {
	// 	logger.Error().Err(err).Msg("Error getting user by ID")
	// 	return
	// }

	// fmt.Println("Имя пользователя:", user.UserName)
	// fmt.Println("ID пользователя:", user.UserID)

	// // nuser := repository.User{
	// // 	UserID:    1999,
	// // 	UserName:  "John Hauzer",
	// // 	InviterID: 1002,
	// // 	CreatedAt: time.Now(),
	// // }

	// // err = r.AddUser(context.Background(), &nuser)
	// // if err != nil {
	// // 	logger.Error().Err(err).Msg("Error getting user by ID")
	// // 	return
	// // }

	// // users, err := r.GetUserByID(context.Background(), 1234)
	// // if err != nil {
	// // 	logger.Error().Err(err).Msg("Error getting user by ID")
	// // 	return
	// // }

	// // fmt.Println("Имя пользователя:", users.UserName)
	// // fmt.Println("ID пользователя:", users.UserID)

	// var usersarray []repository.User
	// var usersarray2 []repository.User

	// usersarray, err = r.GetRefsOfUserFromID(context.Background(), 1002)

	// // fmt.Println(usersarray[0].UserID)
	// for i := range usersarray {
	// 	fmt.Printf("User %d %s %d %s\n", usersarray[i].UserID, usersarray[i].UserName, usersarray[i].InviterID, usersarray[i].CreatedAt)
	// }

	// fmt.Println(r.CountOfAllUsers(context.Background()))
	// fmt.Println(r.CountRefsOfUserFromID(context.Background(), 1002))

	// usersarray2, err = r.GetTopOfRefs(context.Background(), 100)
	// //fmt.Println(usersarray2[0].UserID)
	// for i := range usersarray2 {
	// 	fmt.Printf("User %d %s %d %s\n", usersarray2[i].UserID, usersarray2[i].UserName, usersarray2[i].InviterID, usersarray2[i].CreatedAt)
	// }

	// logger, err := logger.NewLogger()
	// if err != nil {
	// 	log.Fatalf("error init logger: %v", err)
	// }
	// cfg := postgres.Parse()

	// database, err := postgres.New(&cfg)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// r := repository.NewRepository(*database)

	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("error init logger: %v", err)
	}

	cfg, err := config.Parse()
	if err != nil {
		panic(any("cant parse variables from config: " + err.Error()))
	}

	database, err := postgres.New(cfg.Db)
	r := repository.NewRepository(*database)

	if err != nil {
		logger.Error().Err(err).Msg("Postgres start error")
		return
	}
	// defer func() {
	// 	if err = pg.Close(); err != nil {
	// 		logger.Error().Err(err).Msg("error close database")
	// 	}
	// }()

	// repositories

	// services
	userSerivce := service.NewUserService(r)

	platformService := handler.NewPlatformService(
		userSerivce,
	)

	router := chi.NewRouter()
	handlers := route.NewRoutes(logger, router, platformService)
	utils.NewRoutes(
		handlers,
	).Setup()
	chiMux := chi_router.NewChiMux(logger)
	chiMux.Mount("/api/v1", router)

	httpServer := httpserver.New(
		chiMux,
		httpserver.Addr(cfg.MainBackendConfig.Host, cfg.MainBackendConfig.Port),
		httpserver.ReadTimeout(time.Duration(cfg.Timeout)*time.Millisecond),
		httpserver.WriteTimeout(time.Duration(cfg.Timeout)*time.Millisecond),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info().Msg("app - Run - signal: " + s.String())
		logger.Error().Msg(fmt.Sprintf("app - Run - httpServer.Notify: %v", s))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("app - Run - httpServer.Shutdown: %w", err))
	}

	//fmt.Println(cfg.Host)
	logger.Info().Msg("Everything is fine)")
}
