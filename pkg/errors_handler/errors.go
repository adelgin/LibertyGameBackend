package errors_handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var (
	// ErrInternalDatabase Internal database error
	ErrInternalDatabase = fmt.Errorf("internal database error")
	// ErrBadRequest ошибка при невалидных входных параметров
	ErrBadRequest = fmt.Errorf("bad request")
)

// PrepareQueryError ...
type PrepareQueryError struct {
	Entity string
	In     interface{}
}

func (e PrepareQueryError) Error() string {
	return fmt.Sprintf("can't prepare date for %s (%e)", e.Entity, e.In)
}

func NewPrepareQueryError(entity string, in interface{}) *PrepareQueryError {
	return &PrepareQueryError{Entity: entity, In: in}
}

// ValidationError ...
type ValidationError struct {
	Text string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("request validation error: %s", e.Text)
}

func NewValidationError(text string) *ValidationError {
	return &ValidationError{Text: text}
}

type ErrorResponse struct {
	Error          string `json:"error"`
	LocalisedError string `json:"localized_error"`
}

func JError(w http.ResponseWriter, err error) error {
	code := http.StatusInternalServerError
	localizedError := "Внутренняя ошибка!"

	switch err {
	case ErrInternalDatabase:
		code = http.StatusInternalServerError
		localizedError = "Внутренняя ошибка базы данных!"
	case ErrBadRequest:
		code = http.StatusBadRequest
		localizedError = "Неверно указаны входные параметры"

	}

	var validationErr *ValidationError
	if errors.Is(err, validationErr) {
		code = http.StatusBadRequest
		localizedError = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status", strconv.Itoa(code))
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(
		ErrorResponse{
			Error:          err.Error(),
			LocalisedError: localizedError,
		}); err != nil {
		return fmt.Errorf("cannot write response: %w", err)
	}

	return nil
}

func CustomError(w http.ResponseWriter, resp *http.Response) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status", strconv.Itoa(resp.StatusCode))
	w.WriteHeader(resp.StatusCode)
	buf := new(bytes.Buffer)

	_, _ = buf.ReadFrom(resp.Body)
	_, _ = w.Write(buf.Bytes())

	return nil
}
