package handler

import (
	"net/http"
	"strconv"

	"libertyGame/internal/utils"

	"libertyGame/pkg/errors_handler"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
)

// GetUserByID
// @Summary Получение данных о юзере
// @Tags UserService
// @Description Метод возвращает данные юзера, какого именно юзера определяется по id.
// @Accept json
// @Produce json
// @Param id path int false "id"
// @Success 200 {object} repository.User
// @Failure 401 {integer} integer
// @Failure 500 {object} errors_handler.ErrorResponse
// @Router /v1/user/{id} [get]
func (i *Implementation) GetUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
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
