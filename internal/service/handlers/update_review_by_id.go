package handlers

import (
	"ReviewInterfaceAPI/internal/models"
	"ReviewInterfaceAPI/internal/service"
	"ReviewInterfaceAPI/internal/service/helpers"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type UpdateHandler struct {
	Service *service.ReviewService
}

func (h *Handler) UpdateReviewById(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(chi.URLParam(r, "product_id"))
	if err != nil {
		helpers.SendApiError(w, models.ErrInvalidInput)
		return
	}
	reviewID, err := strconv.Atoi(chi.URLParam(r, "review_id"))
	if err != nil {
		helpers.SendApiError(w, models.ErrInvalidInput)
		return
	}
	// Decoding
	var req models.ReviewUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.SendApiError(w, models.ErrInvalidInput)
		return
	}
	// Validation
	updateData := req.Data.Attributes
	validate := validator.New()
	if err := validate.Struct(updateData); err != nil {
		helpers.SendApiError(w, models.ErrInvalidInput)
		return
	}

	reviewService := service.NewReviewService(h.DB)
	updatedReviewID, err := reviewService.UpdateReview(productId, reviewID, updateData)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.SendApiError(w, models.ErrReviewNotFound)
			return
		}
		helpers.SendApiError(w, models.ErrDatabaseProblem)
		return
	}

	successResp := SuccessResponse{}
	successResp.Data.Type = "review"
	successResp.Data.ID = updatedReviewID
	successResp.Data.Attributes = map[string]interface{}{
		"message": fmt.Sprintf("Review with ID %d for product %d updated successfully", updatedReviewID, productId),
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResp)
}
