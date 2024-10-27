package handler

import (
	"net/http"
	"strconv"

	"libertyGame/internal/utils"

	"libertyGame/pkg/errors_handler"

	"github.com/go-chi/chi"

	"github.com/rs/zerolog/log"
)

// GetRefsOfUserFromID
// @Summary Получение данных о юзере
// @Tags UserService
// @Description Метод возвращает список рефералов пользователя, какого именно юзера определяется по id.
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} []repository.User
// @Failure 401 {integer} integer
// @Failure 500 {object} errors_handler.ErrorResponse
// @Router /v1/user/{id}/refs [get]
func (i *Implementation) GetRefsOfUserFromID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			i.SendErrorMessage(err, errors_handler.ErrBadRequest, w)
			return
		}

		refs, err := i.UserService.GetRefsOfUserFromID(r.Context(), userID)
		if err != nil {
			i.SendErrorMessage(err, errors_handler.ErrInternalDatabase, w)
			return
		}

		if err := utils.Json(w, http.StatusOK, refs); err != nil {
			log.Error().Err(err).Msg("Error: ")
		}
	}
}
