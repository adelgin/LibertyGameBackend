package handler

import (
	"net/http"
	"strconv"

	"libertyGame/internal/utils"

	"libertyGame/pkg/errors_handler"

	"github.com/rs/zerolog/log"
)

// GetUserByID
func (i *Implementation) GetUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userIDStr := r.URL.Query().Get("id")
		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			i.SendErrorMessage(err, errors_handler.ErrBadRequest, w)
			return
		}

		user, err := i.UserService.GetUserByID(r.Context(), userID)
		if err != nil {
			i.SendErrorMessage(err, errors_handler.ErrInternalDatabase, w)
			return
		}

		if err := utils.Json(w, http.StatusOK, user); err != nil {
			log.Error().Err(err).Msg("Error: ")
		}
	}
}
