package routes

import (
    "cinetrack/handlers"
    "github.com/gin-gonic/gin"
	"cinetrack/middlewares"
)

func SetupRoutes(r *gin.Engine) {
    // Public routes
    r.GET("/movies", handlers.GetAllMovies)
    r.GET("/movies/:id", handlers.GetMovieByID)
    r.GET("/movies/:id/reviews", handlers.GetReviewsByMovie)

    // Protected routes
    protected := r.Group("/")
    protected.Use(middlewares.AuthMiddleware)
    {
        protected.POST("/movies", handlers.CreateMovie)
        protected.PUT("/movies/:id", handlers.UpdateMovie)
        protected.DELETE("/movies/:id", handlers.DeleteMovie)

        protected.GET("/users", handlers.GetAllUsers)
        protected.GET("/users/:id", handlers.GetUserByID)
        protected.POST("/users", handlers.CreateUser)
        protected.DELETE("/users/:id", handlers.DeleteUser)

        protected.POST("/reviews", handlers.CreateReview)
    }
}