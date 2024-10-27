package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"libertyGame/internal/repository"
	"libertyGame/internal/utils"
	"libertyGame/pkg/errors_handler"

	"github.com/rs/zerolog/log"
)

// UserRegistration
func (i *Implementation) UserRegistration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userRequest repository.User

		if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
			i.SendErrorMessage(err, errors_handler.ErrBadRequest, w)
			return
		}

		userRequest.CreatedAt = time.Now()

		err := i.UserService.AddUser(r.Context(), &userRequest)
		if err != nil {
			i.SendErrorMessage(err, errors_handler.ErrInternalDatabase, w)
			return
		}

		if err := utils.Json(w, http.StatusCreated, userRequest); err != nil {
			log.Error().Err(err).Msg("Error: ")
		}
	}
}
