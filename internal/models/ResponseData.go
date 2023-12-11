package models

type ResponseData struct {
	Type       string `json:"type"`
	ID         int    `json:"id"`
	Attributes Review `json:"attributes"`
}
