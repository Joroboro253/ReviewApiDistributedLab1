package requests

import "github.com/Masterminds/squirrel"

func (s *ReviewService) DeleteReviewsByProductID(productID int) error {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete("").
		From("reviews").
		Where(squirrel.Eq{"product_id": productID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = s.DB.Exec(query, args...)
	return err
}
