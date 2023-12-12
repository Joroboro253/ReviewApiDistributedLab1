package handlers

import (
	"encoding/json"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	helpers "github.com/Joroboro253/ReviewApiDistributedLab/internal/service/heplers"
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

func (h *Handler) CreateReview(w http.ResponseWriter, r *http.Request) {
	productIDStr := chi.URLParam(r, "product_id")
	productId, err := strconv.Atoi(productIDStr)
	if err != nil {
		helpers.SendApiError(w, models.ErrInvalidInput)
		return
	}
	// Decoding
	var reqBody models.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		helpers.SendApiError(w, models.ErrInvalidInput)
		return
	}
	// checking data type
	if reqBody.Data.Type != "review" {
		helpers.SendApiError(w, models.ErrInvalidInput)
		return
	}
	review := reqBody.Data.Attributes
	review.ProductID = productId
	review.CreatedAt = time.Now()
	review.UpdatedAt = time.Now()
	// Validation
	validate := validator.New()
	if err := validate.Struct(review); err != nil {
		helpers.SendApiError(w, models.ErrDatabaseProblem)
		return
	}

	reviewService := requests.NewReviewService(h.DB)
	//reviewID, err := reviewService.CreateReview(&review)
	reviewID, err := reviewService.CreateReview(&review)
	if err != nil {
		log.Printf("Error inserting review into database: %v", err)
		helpers.SendApiError(w, models.ErrDatabaseProblem)
		return
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
	json.NewEncoder(w).Encode(respBody)
}
