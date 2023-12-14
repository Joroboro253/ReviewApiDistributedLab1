package handlers

import (
	_ "github.com/asaskevich/govalidator"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	DB *sqlx.DB
}
