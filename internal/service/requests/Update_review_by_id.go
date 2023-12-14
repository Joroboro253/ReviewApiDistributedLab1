package requests

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/Masterminds/squirrel"
)

func (s *ReviewService) UpdateReview(productId, reviewId int, updateData models.ReviewUpdate) (int, error) {
	// Initialization of SQL-builder queries
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Update("reviews")

	// Add Update Conditions if provided
	if updateData.UserID != nil {
		builder = builder.Set("user_id", *updateData.UserID)
	}
	if updateData.Content != nil {
		builder = builder.Set("content", *updateData.Content)
	}

	query, args, err := builder.Where(squirrel.Eq{"id": reviewId, "product_id": productId}).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	// Executing the query
	var updatedReviewID int
	err = s.DB.QueryRow(query, args...).Scan(&updatedReviewID)
	if err != nil {
		return 0, err
	}

	return updatedReviewID, nil
}
