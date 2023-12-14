package handlers

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	helpers "github.com/Joroboro253/ReviewApiDistributedLab/internal/service/heplers"
	"net/http"
)

func ErrorHandler(handler func(http.ResponseWriter, *http.Request) *models.APIError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if apiErr := handler(w, r); apiErr != nil {
			helpers.SendApiError(w, apiErr)
		}

	}
}
