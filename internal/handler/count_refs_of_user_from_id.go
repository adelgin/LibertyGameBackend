package handler

import (
	"net/http"
	"strconv"

	"libertyGame/internal/utils"

	"libertyGame/pkg/errors_handler"

	"github.com/go-chi/chi"

	"github.com/rs/zerolog/log"
)

// CountRefsOfUserFromID
// @Summary Получение данных о юзере
// @Tags UserService
// @Description Метод возвращает количество рефералов пользователя, какого именно юзера определяется по id.
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} int64
// @Failure 401 {integer} integer
// @Failure 500 {object} errors_handler.ErrorResponse
// @Router /v1/user/{id}/refscount [get]
func (i *Implementation) CountRefsOfUserFromID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			i.SendErrorMessage(err, errors_handler.ErrBadRequest, w)
			return
		}

		refCount, err := i.UserService.CountRefsOfUserFromID(r.Context(), userID)
		if err != nil {
			i.SendErrorMessage(err, errors_handler.ErrInternalDatabase, w)
			return
		}

		if err := utils.Json(w, http.StatusOK, refCount); err != nil {
			log.Error().Err(err).Msg("Error: ")
		}
	}
}
