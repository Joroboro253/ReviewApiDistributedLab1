package handlers

import (
	"encoding/json"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/heplers"
	"log"
	"net/http"
)

func ErrorHandler(handler func(http.ResponseWriter, *http.Request) *models.APIError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if apiErr := handler(w, r); apiErr != nil {
			helpers.SendApiError(w, apiErr)
		}

	}
}

func SendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("Error encoding JSON: %v", err)
		}
	}
}
