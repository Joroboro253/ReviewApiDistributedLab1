package helpers

import (
	"encoding/json"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/requests"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
)

type GetHandler struct {
	DB *sqlx.DB
}

func (h *GetHandler) QueryParameterProcessing(r *http.Request) (string, string, string) {
	sortField := r.URL.Query().Get("sort")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	return sortField, pageStr, limitStr
}

func (h *GetHandler) Pagination(productID int, sortField, pageStr, limitStr string) (map[string]int, map[string]interface{}, *models.APIError) {
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
		return nil, nil, models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
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

	return paginationMeta, response, nil
}

func (h *GetHandler) GenerateResponse(w http.ResponseWriter, response map[string]interface{}) *models.APIError {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return models.NewAPIError(http.StatusInternalServerError, "StatusInternalServerError", "Error encoding response")
	}
	return nil
}
