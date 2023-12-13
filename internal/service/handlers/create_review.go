package handlers

import (
	"encoding/json"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/requests"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CreateHandler struct {
	Service *requests.ReviewService
}

func (h *Handler) CreateReview(w http.ResponseWriter, r *http.Request) *models.APIError {
	productIDStr := chi.URLParam(r, "product_id")
	productId, err := strconv.Atoi(productIDStr)
	if err != nil {
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Wrong format product_id")
	}
	// Decoding
	var reqBody models.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error during JSON decoding")
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
	validate := validator.New()
	if err := validate.Struct(review); err != nil {
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Data validation error")
	}

	reviewService := requests.NewReviewService(h.DB)
	reviewID, err := reviewService.CreateReview(&review)
	if err != nil {
		log.Printf("Error inserting review into database: %v", err)
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error inserting revocation into database")
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
		return models.NewAPIError(http.StatusInternalServerError, "StatusInternalServerError", "Response generation error")
	}
	return nil
}
