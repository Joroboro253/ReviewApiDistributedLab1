package helpers

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"log"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error, statusCode int, errorCode, errorMessage string) *models.APIError {
	log.Printf("Error: %v", err)
	return models.NewAPIError(statusCode, errorCode, errorMessage)
}
