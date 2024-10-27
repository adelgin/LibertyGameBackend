package handler

import (
	"net/http"

	"libertyGame/internal/utils"

	"libertyGame/pkg/errors_handler"

	"github.com/rs/zerolog/log"
)

// CountOfAllUsers
// @Summary Количество юзеров
// @Tags UserService
// @Description Метод возвращает количество юзеров
// @Accept json
// @Produce json
// @Success 200 {object} int64
// @Failure 401 {integer} integer
// @Failure 500 {object} errors_handler.ErrorResponse
// @Router /v1/users/all [get]
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
