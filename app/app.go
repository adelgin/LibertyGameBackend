package app

import (
	"fmt"
	"libertyGame/internal"
	"libertyGame/pkg/logger"
	"log"
)

func Run() {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("error init logger: %v", err)
	}
	cfg := postgres.Parse()

	// r := repository{

	// }

	//user, err := r.GetUserByID(context.Background(), 1001)
	// if err != nil {
	// 	// Обработка ошибки
	// 	logger.Error().Err(err).Msg("Error getting user by ID")
	// 	return // Или выполните другое действие при возникновении ошибки
	// }

	// // Доступ к данным пользователя
	// fmt.Println("Имя пользователя:", user.FirstName)
	// fmt.Println("ID пользователя:", user.ID)

	fmt.Println(cfg.Host)
	logger.Info().Msg("Everything is fine)")
}
