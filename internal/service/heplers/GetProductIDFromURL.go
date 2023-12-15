package helpers

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func GetProductIDFromURL(r *http.Request) (int, *models.APIError) {
	productIDStr := chi.URLParam(r, "product_id")
	productId, err := strconv.Atoi(productIDStr)
	if err != nil {
		errorMsg := "Wrong format product_id"
		log.Printf("%s: %v", errorMsg, err)
		return 0, models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
	}
	return productId, nil
}
