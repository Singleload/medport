package handlers

import (
	"net/http"

	"github.com/svenskhalsovard/api/internal/dto"
	"github.com/svenskhalsovard/api/internal/version"
)

// HealthCheck handles health check requests
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := dto.NewHealthCheckResponse(version.Version)
	RespondJSON(w, http.StatusOK, dto.NewSuccessResponse(response))
}

// Version handles version requests
func Version(w http.ResponseWriter, r *http.Request) {
	response := dto.VersionResponse{
		Version:   version.Version,
		BuildTime: version.BuildTime,
		GitCommit: version.GitCommit,
	}
	RespondJSON(w, http.StatusOK, dto.NewSuccessResponse(response))
}