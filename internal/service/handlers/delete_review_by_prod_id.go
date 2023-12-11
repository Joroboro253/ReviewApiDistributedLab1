package handlers

import (
	"ReviewInterfaceAPI/internal/models"
	"ReviewInterfaceAPI/internal/service"
	"ReviewInterfaceAPI/internal/service/helpers"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type DeleteHandler struct {
	Service *service.ReviewService
}

func (h *Handler) DeleteReviews(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(chi.URLParam(r, "product_id"))
	if err != nil {
		helpers.SendApiError(w, models.ErrInvalidInput)
		return
	}
	reviewService := service.NewReviewService(h.DB)
	err = reviewService.DeleteReviewsByProductID(productID)
	if err != nil {
		helpers.SendApiError(w, models.ErrDatabaseProblem)
		return
	}
	// response about successful deleting
	w.WriteHeader(http.StatusNoContent)
}
