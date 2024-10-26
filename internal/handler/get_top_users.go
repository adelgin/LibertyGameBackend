package handler

import (
	"net/http"
	"strconv"

	"libertyGame/internal/utils"

	"libertyGame/pkg/errors_handler"

	"github.com/rs/zerolog/log"
)

// GetTopUsers
func (i *Implementation) GetTopOfRefs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		count, err := strconv.Atoi(r.URL.Query().Get("count"))
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
