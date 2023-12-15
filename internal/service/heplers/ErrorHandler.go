package helpers

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"net/http"
)

func ErrorHandler(handler func(http.ResponseWriter, *http.Request) *models.APIError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if apiErr := handler(w, r); apiErr != nil {
			SendApiError(w, apiErr)
		}
	}
}
