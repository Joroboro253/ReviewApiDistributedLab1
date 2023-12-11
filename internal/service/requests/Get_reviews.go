package requests

import (
	"fmt"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/models"
	"github.com/Masterminds/squirrel"
	"log"
	"math"
)

func (s *ReviewService) GetReviewsByProductID(productID int, sortField string, page, limit int) ([]models.Review, int, int, error) {
	// Pagination param check
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10 // Default setting
	}
	// Getting revews
	countBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("COUNT(*)").
		From("reviews").
		Where(squirrel.Eq{"product_id": productID})

	// Counting the total number of reviews
	countQuery, countArgs, err := countBuilder.ToSql()
	log.Printf("Count Query: %s, Args: %v", countQuery, countArgs)
	if err != nil {
		log.Printf("error building count SQL query: %v", err)
		return nil, 0, 0, fmt.Errorf("error building count SQL query: %w", err)
	}

	var totalReviews int
	err = s.DB.Get(&totalReviews, countQuery, countArgs...)
	if err != nil {
		log.Printf("error executing count SQL query: %v", err)
		return nil, 0, 0, fmt.Errorf("error executing count SQL query: %w", err)
	}

	reviewBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("*").
		From("reviews").
		Where(squirrel.Eq{"product_id": productID})

	if sortField != "" {
		reviewBuilder = reviewBuilder.OrderBy(sortField)
	}

	query, args, err := reviewBuilder.Limit(uint64(limit)).Offset(uint64((page - 1) * limit)).ToSql()
	if err != nil {
		log.Printf("error building SQL query: %v", err)
		return nil, 0, 0, fmt.Errorf("error building SQL query: %w", err)
	}
	// Pagination
	var reviews []models.Review
	err = s.DB.Select(&reviews, query, args...)
	if err != nil {
		log.Printf("error executing SQL query: %v", err)
		return nil, 0, 0, fmt.Errorf("error executing SQL query: %w", err)
	}

	// Calculation of total pages
	totalPages := int(math.Ceil(float64(totalReviews) / float64(limit)))

	// Return of results
	return reviews, totalReviews, totalPages, nil
}
