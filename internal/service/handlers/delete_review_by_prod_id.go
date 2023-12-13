package handlers

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/requests"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type DeleteHandler struct {
	Service *requests.ReviewService
}

func (h *Handler) DeleteReviews(w http.ResponseWriter, r *http.Request) *models.APIError {
	productID, err := strconv.Atoi(chi.URLParam(r, "product_id"))
	if err != nil {
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Wrong format product_id")
	}
	reviewService := requests.NewReviewService(h.DB)
	err = reviewService.DeleteReviewsByProductID(productID)
	if err != nil {
		return models.NewAPIError(http.StatusBadGateway, "StatusBadGateway", "Database problem")
	}
	// response about successful deleting
	w.WriteHeader(http.StatusNoContent)
	return nil
}
