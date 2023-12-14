package handlers

import (
	"encoding/json"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	helpers "github.com/Joroboro253/ReviewApiDistributedLab/internal/service/heplers"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/requests"
	"log"
	"net/http"
	"time"
)

type CreateHandler struct {
	Service *requests.ReviewService
}

func (h *Handler) CreateReview(w http.ResponseWriter, r *http.Request) *models.APIError {
	productId, err := helpers.GetProductIDFromURL(r, w)
	if err != nil {
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Wrong format product_id")
	}
	// Decoding
	var reqBody models.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		errorMsg := "Error during JSON decoding"
		log.Printf("%s: %v", errorMsg, err)
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
	}
	// checking data type
	if reqBody.Data.Type != "review" {
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Incorrect data type")
	}
	review := reqBody.Data.Attributes
	review.ProductID = productId
	review.CreatedAt = time.Now()
	review.UpdatedAt = time.Now()

	// Validation
	if validationErr := helpers.ValidateReviewAttributes(&review); validationErr != nil {
		return validationErr
	}

	reviewService := requests.NewReviewService(h.DB)
	reviewID, err := reviewService.CreateReview(&review)
	if err != nil {
		log.Printf("Error inserting review into database: %v", err)
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error inserting review into database")
	}
	review.ID = reviewID

	// Query generation
	respBody := models.ResponseBody{
		Data: models.ResponseData{
			Type:       "review",
			ID:         review.ID,
			Attributes: review,
		},
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		log.Printf("Response generation error: %v", err)
		return models.NewAPIError(http.StatusInternalServerError, "StatusInternalServerError", "Response generation error")
	}
	return nil
}
