package routes

import (
    "cinetrack/handlers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    // Movies
    r.GET("/movies", handlers.GetAllMovies)
    r.GET("/movies/:id", handlers.GetMovieByID)
    r.POST("/movies", handlers.CreateMovie)
    r.PUT("/movies/:id", handlers.UpdateMovie)
    r.DELETE("/movies/:id", handlers.DeleteMovie)

    // Users
    r.GET("/users", handlers.GetAllUsers)
    r.GET("/users/:id", handlers.GetUserByID)
    r.POST("/users", handlers.CreateUser)
    r.DELETE("/users/:id", handlers.DeleteUser)

    // Reviews
    r.POST("/reviews", handlers.CreateReview)
    r.GET("/movies/:id/reviews", handlers.GetReviewsByMovie)
}