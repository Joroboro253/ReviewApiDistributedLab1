package requests

import (
	"ReviewInterfaceAPI/internal/models"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"log"
)

type ReviewService struct {
	DB *sqlx.DB
}

func NewReviewService(db *sqlx.DB) *ReviewService {
	return &ReviewService{DB: db}
}

func (s *ReviewService) CreateReview(review *models.Review) (int, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.Insert("reviews").
		Columns("product_id", "user_id", "rating", "content", "created_at", "updated_at").
		Values(review.ProductID, review.UserID, review.Rating, review.Content, review.CreatedAt, review.UpdatedAt).
		Suffix("RETURNING id").
		ToSql()
	log.Printf("Executing SQL query: %s with args: %v", query, args)
	if err != nil {
		log.Printf("error building insert SQL query: %v", err)
		return 0, fmt.Errorf("error building insert SQL query: %w", err)
	}

	var reviewID int
	err = s.DB.QueryRow(query, args...).Scan(&reviewID)
	if err != nil {
		log.Printf("error executing insert SQL query: %v", err)
		return 0, fmt.Errorf("error executing insert SQL query: %w", err)
	}

	return reviewID, nil
}
