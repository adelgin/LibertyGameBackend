package handler

import (
	"net/http"

	"libertyGame/internal/service"
	"libertyGame/pkg/errors_handler"

	"github.com/rs/zerolog/log"
)

// HandleInterface interface for all handlers
type HandleInterface interface {
	GetUserByID() http.HandlerFunc
	CountOfAllUsers() http.HandlerFunc
	GetRefsOfUserFromID() http.HandlerFunc
	CountRefsOfUserFromID() http.HandlerFunc
	GetTopOfRefs() http.HandlerFunc
}

type Implementation struct {
	UserService service.UserService
	HandleInterface
}

func NewPlatformService(
	userService service.UserService,
) *Implementation {
	return &Implementation{
		UserService: userService,
	}
}

func (i *Implementation) SendErrorMessage(err, httpError error, w http.ResponseWriter) {
	log.Error().Err(err).Msg("Error: ")
	if err := errors_handler.JError(w, httpError); err != nil {
		log.Error().Msg("Cannot send error message")
	}
}
