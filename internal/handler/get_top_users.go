package handler

import (
	"net/http"
	"strconv"

	"libertyGame/internal/utils"

	"libertyGame/pkg/errors_handler"

	"github.com/go-chi/chi"

	"github.com/rs/zerolog/log"
)

// GetTopOfRefs
// @Summary Получение топов игроков
// @Tags UserService
// @Description Метод возвращает топы игроков, количество людей в топе определяется по count.
// @Accept json
// @Produce json
// @Param count path int true "count"
// @Success 200 {object} []repository.Top_User
// @Failure 401 {integer} integer
// @Failure 500 {object} errors_handler.ErrorResponse
// @Router /v1/users/{count}/top [get]
func (i *Implementation) GetTopOfRefs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		count, err := strconv.ParseInt(chi.URLParam(r, "count"), 10, 64)
		if err != nil || count < 0 {
			i.SendErrorMessage(err, errors_handler.ErrBadRequest, w)
			return
		}

		res, err := i.UserService.GetTopOfRefs(r.Context(), count)
		if err != nil {
			i.SendErrorMessage(err, errors_handler.ErrInternalDatabase, w)
			return
		}

		if err := utils.Json(w, http.StatusOK, res); err != nil {
			log.Error().Err(err).Msg("Error: ")
		}
	}
}
