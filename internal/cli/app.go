package cli

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"net/http"
)

// Delete App later
type App struct {
	Router *chi.Mux
	DB     *sqlx.DB
}

func NewApp(db *sqlx.DB) *App {
	app := &App{
		Router: chi.NewRouter(),
		DB:     db,
	}
	app.setupRoutes()
	return app
}

func (app *App) Start(addr string) error {
	return http.ListenAndServe(addr, app.Router)
}

func (app *App) setupRoutes() {
	reviewHandler := &handlers.Handler{
		DB: app.DB,
	}
	app.Router.Route("/products/{product_id}/reviews", func(r chi.Router) {
		r.Post("/", handlers.ErrorHandler(reviewHandler.CreateReview))
		r.Get("/", handlers.ErrorHandler(reviewHandler.GetReviews))
		r.Delete("/", handlers.ErrorHandler(reviewHandler.DeleteReviews))
		r.Patch("/{review_id}", handlers.ErrorHandler(reviewHandler.UpdateReviewById))
	},
	)

}
