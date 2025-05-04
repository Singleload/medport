package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/svenskhalsovard/api/internal/dto"
)

var validate = validator.New()

// RespondJSON sends a JSON response
func RespondJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal JSON response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// RespondError sends an error response
func RespondError(w http.ResponseWriter, err error) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		details := formatValidationErrors(validationErrors)
		errResp := dto.NewErrorResponse(dto.ErrorCodeValidationFailed, "Validation failed", details)
		RespondJSON(w, dto.GetStatusCodeForErrorCode(dto.ErrorCodeValidationFailed), errResp)
		return
	}

	// Default to internal server error
	errResp := dto.NewErrorResponse(dto.ErrorCodeInternalServerError, err.Error(), nil)
	RespondJSON(w, dto.GetStatusCodeForErrorCode(dto.ErrorCodeInternalServerError), errResp)
}

// formatValidationErrors formats validation errors into a user-friendly format
func formatValidationErrors(errs validator.ValidationErrors) []dto.ValidationError {
	validationErrors := make([]dto.ValidationError, 0, len(errs))
	for _, err := range errs {
		validationErrors = append(validationErrors, dto.ValidationError{
			Field:   err.Field(),
			Message: getValidationErrorMessage(err),
		})
	}
	return validationErrors
}

// getValidationErrorMessage returns a user-friendly error message for a validation error
func getValidationErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Must be a valid email address"
	case "min":
		if err.Type().Kind().String() == "string" {
			return fmt.Sprintf("Must be at least %s characters", err.Param())
		}
		return fmt.Sprintf("Must be at least %s", err.Param())
	case "max":
		if err.Type().Kind().String() == "string" {
			return fmt.Sprintf("Must be at most %s characters", err.Param())
		}
		return fmt.Sprintf("Must be at most %s", err.Param())
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", err.Param())
	default:
		return fmt.Sprintf("Failed validation on %s", err.Tag())
	}
}

// ParseJSON parses a JSON request body
func ParseJSON(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %w", err)
	}
	defer r.Body.Close()

	if len(body) == 0 {
		return errors.New("empty request body")
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	return validate.Struct(v)
}

// ParseIDParam parses an ID parameter from the request URL
func ParseIDParam(r *http.Request, paramName string) (int64, error) {
	idStr := chi.URLParam(r, paramName)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s parameter: %w", paramName, err)
	}
	return id, nil
}