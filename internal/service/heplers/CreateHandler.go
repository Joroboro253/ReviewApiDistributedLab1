package helpers

import (
	"encoding/json"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/requests"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

type CreateHandler struct {
	DB *sqlx.DB
}

func (h *CreateHandler) CreateReview(productId int, reqBody models.UpdateRequest) (models.Review, *models.APIError) {
	review := reqBody.Data.Attributes
	review.ProductID = productId
	review.CreatedAt = time.Now()
	review.UpdatedAt = time.Now()

	if validationErr := ValidateReviewAttributes(&review); validationErr != nil {
		return models.Review{}, validationErr
	}

	reviewService := requests.NewReviewService(h.DB)
	reviewID, err := reviewService.CreateReview(&review)
	if err != nil {
		log.Printf("Error inserting review into database: %v", err)
		return models.Review{}, models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error inserting review into database")
	}
	review.ID = reviewID
	return review, nil
}

func (h *CreateHandler) DecodeRequestBody(r *http.Request) (models.UpdateRequest, *models.APIError) {
	var reqBody models.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		errorMsg := "Error during JSON decoding"
		log.Printf("%s: %v", errorMsg, err)
		return models.UpdateRequest{}, models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
	}
	return reqBody, nil
}

func (h *CreateHandler) GenerateResponse(w http.ResponseWriter) *models.APIError {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusCreated)
	return nil
}
