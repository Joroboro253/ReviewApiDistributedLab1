package models

type ReviewAttributes struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Rating  int    `json:"rating" validate:"required,min=1,max=5"`
}
