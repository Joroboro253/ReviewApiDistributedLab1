package handlers

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	helpers "github.com/Joroboro253/ReviewApiDistributedLab/internal/service/heplers"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/requests"
	"net/http"
)

type CreateHandler struct {
	Service       *requests.ReviewService
	ReviewHandler helpers.ReviewHandlerInterface
}

func (h *Handler) CreateReview(w http.ResponseWriter, r *http.Request) *models.APIError {
	handler := helpers.CreateHandler{
		DB: h.DB,
	}
	productId, apiErr := handler.GetProductIDFromURL(r)
	if apiErr != nil {
		return apiErr
	}
	// Decoding
	reqBody, apiErr := handler.DecodeRequestBody(r)
	if apiErr != nil {
		return apiErr
	}
	// checking data type
	if reqBody.Data.Type != "review" {
		return models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Incorrect data type")
	}
	review, apiErr := handler.CreateReview(productId, reqBody)
	if apiErr != nil {
		return apiErr
	}
	return handler.GenerateResponse(w, review)
}
