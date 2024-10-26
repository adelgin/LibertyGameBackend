package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"libertyGame/pkg/errors_handler"

	"github.com/pkg/errors"
	"github.com/ydb-platform/ydb-go-sdk/v3/log"
)

func Json(w http.ResponseWriter, code int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status", strconv.Itoa(code))
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		return fmt.Errorf("cannot write response: %w", err)
	}

	return nil
}

func HandleError(w http.ResponseWriter, err error, customError error) {
	log.Error(err)

	if err := errors_handler.JError(w, customError); err != nil {
		log.Error(errors.Wrap(err, "Error sending error response"))
	}
}
