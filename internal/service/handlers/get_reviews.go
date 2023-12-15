package handlers

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	helpers "github.com/Joroboro253/ReviewApiDistributedLab/internal/service/heplers"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/requests"
	"log"
	"net/http"
	"strconv"
)

type GetHandler struct {
	Service       *requests.ReviewService
	ReviewHandler helpers.ReviewHandlerInterface
}

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) *models.APIError {
	handler := helpers.GetHandler{
		DB: h.DB,
	}
	productID, apiErr := helpers.GetProductIDFromURL(r)
	if apiErr != nil {
		errorMsg := "Error getting product ID from URL"
		log.Printf("%s: %v", errorMsg, apiErr)
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
	}
	if productID < 1 {
		errorMsg := "Invalid product_id"
		log.Printf("%s: %d", errorMsg, productID)
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
	}
	sortField, pageStr, limitStr := handler.QueryParameterProcessing(r)
	// Conversion page and limit
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10 // default value
	}
	// Pagination
	_, response, apiErr := handler.Pagination(productID, sortField, strconv.Itoa(page), strconv.Itoa(limit))
	if apiErr != nil {
		return apiErr
	}

	// Generate and send response
	return handler.GenerateResponse(w, response)
}
