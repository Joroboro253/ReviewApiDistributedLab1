package requests

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/Masterminds/squirrel"
	"log"
	"net/http"
)

func (s *ReviewService) CreateReview(review *models.Review) (int, *models.APIError) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.Insert("reviews").
		Columns("product_id", "user_id", "content", "created_at", "updated_at").
		Values(review.ProductID, review.UserID, review.Content, review.CreatedAt, review.UpdatedAt).
		Suffix("RETURNING id").
		ToSql()
	log.Printf("Executing SQL query: %s with args: %v", query, args)
	if err != nil {
		log.Printf("error building insert SQL query: %v", err)
		return 0, models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error inserting review into database")

		//return 0, fmt.Errorf("error building insert SQL query: %w", err)
	}

	var reviewID int
	err = s.DB.QueryRow(query, args...).Scan(&reviewID)
	if err != nil {
		log.Printf("error executing insert SQL query: %v", err)
		return 0, models.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error inserting review into database")
	}

	return reviewID, nil
}
