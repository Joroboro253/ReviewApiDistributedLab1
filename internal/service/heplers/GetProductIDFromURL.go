package helpers

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func GetProductIDFromURL(r *http.Request, w http.ResponseWriter) (int, *models.APIError) {
	productIDStr := chi.URLParam(r, "product_id")
	productId, err := strconv.Atoi(productIDStr)
	if err != nil {
		return 0, HandleError(w, err, http.StatusBadRequest, "StatusBadRequest", "Wrong format product_id")
	}
	return productId, nil
}
