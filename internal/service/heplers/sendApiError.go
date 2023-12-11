package helpers

//import (
//	"ReviewInterfaceAPI/internal/models"
//	"encoding/json"
//	"net/http"
//)
//
//func SendApiError(w http.ResponseWriter, apiErr *models.APIError) {
//	w.Header().Set("Content-Type", "application/vnd.api+json")
//	w.WriteHeader(apiErr.Status)
//	json.NewEncoder(w).Encode(map[string][]models.APIError{"errors": {*apiErr}})
//}
