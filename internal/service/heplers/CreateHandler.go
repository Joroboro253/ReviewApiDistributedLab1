package helpers

import (
	"encoding/json"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/requests"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
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

func (h *CreateHandler) GetProductIDFromURL(r *http.Request) (int, *models.APIError) {
	productIDStr := chi.URLParam(r, "product_id")
	productId, err := strconv.Atoi(productIDStr)
	if err != nil {
		errorMsg := "Wrong format product_id"
		log.Printf("%s: %v", errorMsg, err)
		return 0, models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
	}
	return productId, nil
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

func (h *CreateHandler) GenerateResponse(w http.ResponseWriter, review models.Review) *models.APIError {
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
