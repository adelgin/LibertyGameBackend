package app

import (
	// "fmt"

	"context"
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

	"github.com/go-chi/chi"
	_ "github.com/jackc/pgx/v4/stdlib"

	_ "github.com/lib/pq"
)

func Run() {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("error init logger: %v", err)
	}

	cfg, err := config.Parse()
	if err != nil {
		panic(any("cant parse variables from config: " + err.Error()))
	}

	// Создание соединения с базой данных
	database, err := postgres.New(cfg.Db) // Передаем cfg.Db в postgres.New
	if err != nil {
		logger.Error().Err(err).Msg("Postgres start error")
		return
	}

	// Создание репозитория
	r := repository.NewRepository(database) // Передаем *database в repository.NewRepository

	r.CreateTable(context.Background())

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
		logger.Error().Msg(fmt.Sprintf("app - Run - httpServer.Shutdown: %s", err))
	}
}
