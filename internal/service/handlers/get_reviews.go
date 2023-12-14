package handlers

import (
	"encoding/json"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/requests"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type GetHandler struct {
	Service *requests.ReviewService
}

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) *models.APIError {
	log.Printf("It's get reviews")
	// Extracting product_id from URL
	productIDStr := chi.URLParam(r, "product_id")
	productID, err := strconv.Atoi(productIDStr)
	//productID, err = helpers.GetProductIDFromURL(r, w)

	if err != nil {
		errorMsg := "Wrong format product_id"
		log.Printf("%s: %v", errorMsg, err)
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
	}
	if productID < 1 {
		errorMsg := "Invalid product_id"
		log.Printf("%s: %d", errorMsg, productID)
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
	}
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

	reviewService := requests.NewReviewService(h.DB)
	reviews, totalReviews, totalPages, err := reviewService.GetReviewsByProductID(productID, sortField, page, limit)

	if err != nil {
		errorMsg := "Error retrieving reviews"
		log.Printf("%s: %v", errorMsg, err)
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
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
	// ???
	return nil
}
