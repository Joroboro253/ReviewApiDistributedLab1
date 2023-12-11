package requests

import "github.com/jmoiron/sqlx"

type ReviewService struct {
	DB *sqlx.DB
}

func NewReviewService(db *sqlx.DB) *ReviewService {
	return &ReviewService{DB: db}
}
