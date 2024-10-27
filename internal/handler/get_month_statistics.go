package handler

import (
	"encoding/json"
	"net/http"

	"libertyGame/pkg/errors_handler"

	"github.com/rs/zerolog/log"
)

// GetMonthStatisticsHandler
// @Summary Получение данных о количестве присоединившихся юзеров по месяцам.
// @Tags UserService
// @Description Метод возвращает количество присоединившихся юзеров по месяцам.
// @Accept json
// @Produce json
// @Success 200 {object} []repository.MonthStatistics
// @Failure 401 {integer} integer
// @Failure 500 {object} errors_handler.ErrorResponse
// @Router /v1/monthstat [get]
func (i *Implementation) GetMonthStatisticsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		stats, err := i.UserService.GetMonthStatistics(ctx)
		if err != nil {
			i.SendErrorMessage(err, errors_handler.ErrInternalDatabase, w)
			return
		}

		jsonStats, err := json.Marshal(stats)
		if err != nil {
			i.SendErrorMessage(err, errors_handler.ErrBadRequest, w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonStats)
		if err != nil {
			log.Error().Err(err).Msg("Error while writing JSON-answer")
		}
	}
}
