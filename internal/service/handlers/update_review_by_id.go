package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/heplers"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/requests"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type UpdateHandler struct {
	Service *requests.ReviewService
}

func (h *Handler) UpdateReviewById(w http.ResponseWriter, r *http.Request) *models.APIError {
	productId, err := strconv.Atoi(chi.URLParam(r, "product_id"))
	if err != nil {
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Wrong format product_id")
	}
	reviewID, err := strconv.Atoi(chi.URLParam(r, "review_id"))
	if err != nil {
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error during JSON decoding")
	}
	// Decoding
	var req models.ReviewUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error during JSON decoding")
	}
	var reqBody models.UpdateRequest
	review := reqBody.Data.Attributes
	// Validation
	if validationErr := helpers.ValidateReviewAttributes(&review); validationErr != nil {
		return validationErr
	}

	updateData := req.Data.Attributes
	validate := validator.New()
	if err := validate.Struct(updateData); err != nil {
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Data validation error")
	}

	reviewService := requests.NewReviewService(h.DB)
	updatedReviewID, err := reviewService.UpdateReview(productId, reviewID, updateData)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.NewAPIError(http.StatusNotFound, "StatusNotFound", "Object not found")
		}
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error inserting revocation into database")

	}

	successResp := models.SuccessResponse{}
	successResp.Data.Type = "review"
	successResp.Data.ID = updatedReviewID
	successResp.Data.Attributes = map[string]interface{}{
		"message": fmt.Sprintf("Review with ID %d for product %d updated successfully", updatedReviewID, productId),
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResp)
	return nil
}
