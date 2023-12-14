package helpers

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	//"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/handlers"
	"net/http"
)

type ReviewHandlerInterface interface {
	GetProductIDFromURL(r *http.Request) (int, *models.APIError)
	DecodeRequestBody(r *http.Request) (models.UpdateRequest, *models.APIError)
	GenerateResponse(w http.ResponseWriter, review models.Review) *models.APIError
	CreateReview(productId int, reqBody models.UpdateRequest) (models.Review, *models.APIError)
}
