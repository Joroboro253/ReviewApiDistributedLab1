package models

type SuccessResponse struct {
	Data struct {
		Type       string      `json:"type"`
		ID         int         `json:"id"`
		Attributes interface{} `json:"attributes,omitempty"`
	} `json:"data"`
}
