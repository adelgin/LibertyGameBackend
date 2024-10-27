package handler

import (
	"encoding/json"
	"net/http"

	"libertyGame/pkg/errors_handler"

	"github.com/rs/zerolog/log"
)

// GetMonthStatisticsHandler
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
