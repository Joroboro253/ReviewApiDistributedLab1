package handlers

import (
	"ReviewInterfaceAPI/internal/models"
	"ReviewInterfaceAPI/internal/service"
	"ReviewInterfaceAPI/internal/service/helpers"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type GetHandler struct {
	Service *service.ReviewService
}

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) {
	log.Printf("It`s get reviews")
	// Extracting product_id from URL
	productID, err := strconv.Atoi(chi.URLParam(r, "product_id"))
	if err != nil {
		helpers.SendApiError(w, models.ErrInvalidInput)
		return
	}
	log.Printf("Product id: %v", productID)
	// Query parameter processing
	sortField := r.URL.Query().Get("sort")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// Conversion page and limit
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10 // default value
	}

	reviewService := service.NewReviewService(h.DB)
	reviews, totalReviews, totalPages, err := reviewService.GetReviewsByProductID(productID, sortField, page, limit)
	if err != nil {
		helpers.SendApiError(w, models.ErrDatabaseProblem)
		return
	}

	// Pagination metadata
	paginationMeta := map[string]int{
		"totalReviews": totalReviews,
		"totalPages":   totalPages,
		"currentPage":  page,
		"limit":        limit,
	}

	// Response formation
	response := map[string]interface{}{
		"data": reviews,
		"meta": paginationMeta,
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
