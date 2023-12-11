package handlers

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	helpers "github.com/Joroboro253/ReviewApiDistributedLab/internal/service/heplers"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/requests"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type DeleteHandler struct {
	Service *requests.ReviewService
}

func (h *Handler) DeleteReviews(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(chi.URLParam(r, "product_id"))
	if err != nil {
		helpers.SendApiError(w, models.ErrInvalidInput)
		return
	}
	reviewService := requests.NewReviewService(h.DB)
	err = reviewService.DeleteReviewsByProductID(productID)
	if err != nil {
		helpers.SendApiError(w, models.ErrDatabaseProblem)
		return
	}
	// response about successful deleting
	w.WriteHeader(http.StatusNoContent)
}
