package handler

import (
	"net/http"

	"libertyGame/internal/utils"

	"libertyGame/pkg/errors_handler"

	"github.com/rs/zerolog/log"
)

// CountOfAllUsers
func (i *Implementation) CountOfAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		count, err := i.UserService.CountOfAllUsers(r.Context())
		if err != nil {
			i.SendErrorMessage(err, errors_handler.ErrInternalDatabase, w)
			return
		}

		if err := utils.Json(w, http.StatusOK, count); err != nil {
			log.Error().Err(err).Msg("Error: ")
		}
	}
}
