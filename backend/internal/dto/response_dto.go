package dto

import (
	"net/http"
	"time"
)

// Response represents a generic API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

// ErrorInfo represents error information in an API response
type ErrorInfo struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors represents a collection of validation errors
type ValidationErrors []ValidationError

// NewSuccessResponse creates a new success response
func NewSuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(code string, message string, details interface{}) Response {
	return Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
			Details: details,
		},
	}
}

// HealthCheckResponse represents a health check response
type HealthCheckResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Time    string `json:"time"`
}

// VersionResponse represents a version response
type VersionResponse struct {
	Version   string `json:"version"`
	BuildTime string `json:"buildTime"`
	GitCommit string `json:"gitCommit"`
}

// NewHealthCheckResponse creates a new health check response
func NewHealthCheckResponse(version string) HealthCheckResponse {
	return HealthCheckResponse{
		Status:  "ok",
		Version: version,
		Time:    time.Now().Format(time.RFC3339),
	}
}

// Error codes
const (
	ErrorCodeInvalidRequest      = "INVALID_REQUEST"
	ErrorCodeValidationFailed    = "VALIDATION_FAILED"
	ErrorCodeResourceNotFound    = "RESOURCE_NOT_FOUND"
	ErrorCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrorCodeUnauthorized        = "UNAUTHORIZED"
	ErrorCodeForbidden           = "FORBIDDEN"
	ErrorCodePaymentFailed       = "PAYMENT_FAILED"
	ErrorCodeBookingFailed       = "BOOKING_FAILED"
)

// HTTP status code mapping
var errorCodeToStatusCode = map[string]int{
	ErrorCodeInvalidRequest:      http.StatusBadRequest,
	ErrorCodeValidationFailed:    http.StatusUnprocessableEntity,
	ErrorCodeResourceNotFound:    http.StatusNotFound,
	ErrorCodeInternalServerError: http.StatusInternalServerError,
	ErrorCodeUnauthorized:        http.StatusUnauthorized,
	ErrorCodeForbidden:           http.StatusForbidden,
	ErrorCodePaymentFailed:       http.StatusBadRequest,
	ErrorCodeBookingFailed:       http.StatusBadRequest,
}

// GetStatusCodeForErrorCode returns the HTTP status code for an error code
func GetStatusCodeForErrorCode(code string) int {
	if statusCode, ok := errorCodeToStatusCode[code]; ok {
		return statusCode
	}
	return http.StatusInternalServerError
}