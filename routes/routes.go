package routes

import (
	"cinetrack/handlers"
	"cinetrack/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// --------------------------------------------------
	// PUBLIC ROUTES
	// --------------------------------------------------

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.GET("/movies", handlers.GetAllMovies)

	// Static routes ko :id se pehle rakho
	r.GET("/movies/high-rated", handlers.GetHighRatedMovies)
	r.GET("/movies/recent", handlers.GetLatestMovies)

	r.GET("/movies/:id", handlers.GetMovieByID)
	r.GET("/movies/:id/reviews", handlers.GetReviewsByMovie)

	r.GET("/stats", handlers.GetStats)

	// --------------------------------------------------
	// PROTECTED ROUTES
	// --------------------------------------------------

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware)

	{
		// Current logged-in user
		protected.GET("/me", handlers.GetCurrentUser)

		// User tracked movies
		protected.GET("/my-movies", handlers.GetMyMovies)

		// Movie tracking
		protected.GET(
			"/movies/:id/tracking",
			handlers.GetTracking,
		)

		protected.PUT(
			"/movies/:id/tracking",
			handlers.SaveTracking,
		)

		// Movie management
		protected.POST(
			"/movies",
			handlers.CreateMovie,
		)

		protected.PUT(
			"/movies/:id",
			handlers.UpdateMovie,
		)

		protected.DELETE(
			"/movies/:id",
			handlers.DeleteMovie,
		)

		protected.POST(
			"/movies/batch",
			handlers.BatchCreateMovies,
		)

		protected.DELETE(
			"/movies/:id/permanent",
			handlers.HardDeleteMovie,
		)
	}
}