package helpers

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

func ValidateReviewAttributes(attributes *models.Review) *models.APIError {
	if err := validate.Struct(attributes); err != nil {
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error during JSON decoding")
	}
	return nil
}
