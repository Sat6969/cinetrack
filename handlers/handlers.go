package handlers

import (
	"cinetrack/database"
	"cinetrack/models"
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// --------------------------------------------------
// MOVIES
// --------------------------------------------------

func GetAllMovies(c *gin.Context) {
	ctx := c.Request.Context()

	var movies []models.Movie

	err := database.DB.WithContext(ctx).
		Preload("Genres").
		Find(&movies).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}

	// Frontend direct array expect karta hai
	c.JSON(200, movies)
}

func GetMovieByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	var movie models.Movie

	err := database.DB.WithContext(ctx).
		Preload("Genres").
		Preload("Reviews").
		First(&movie, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"message": "movie not found",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}

	// Frontend direct movie object expect karta hai
	c.JSON(200, movie)
}

func CreateMovie(c *gin.Context) {
	ctx := c.Request.Context()

	var input struct {
		Title       string   `json:"title"`
		Rating      int      `json:"rating"`
		ReleaseYear int      `json:"release_year"`
		Description string   `json:"description"`
		Genres      []string `json:"genres"`
	}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid data",
		})
		return
	}

	movie := models.Movie{
		Title:        input.Title,
		Rating:       input.Rating,
		Release_year: input.ReleaseYear,
		Description:  input.Description,
	}

	for _, name := range input.Genres {
		var genre models.Genre

		result := database.DB.WithContext(ctx).
			Where("name = ?", name).
			FirstOrCreate(
				&genre,
				models.Genre{Name: name},
			)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"message": "failed to create genre",
			})
			return
		}

		movie.Genres = append(movie.Genres, genre)
	}

	err = database.DB.WithContext(ctx).
		Create(&movie).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "could not create movie",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "movie created successfully",
	})
}

func UpdateMovie(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	var movie models.Movie

	err := database.DB.WithContext(ctx).
		First(&movie, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"message": "movie not found",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}

	var updatedMovie models.Movie

	err = c.ShouldBindJSON(&updatedMovie)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid input",
		})
		return
	}

	err = database.DB.WithContext(ctx).
		Model(&movie).
		Updates(&updatedMovie).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "could not update movie",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "movie updated successfully",
	})
}

func DeleteMovie(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	var movie models.Movie

	err := database.DB.WithContext(ctx).
		First(&movie, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"message": "movie not found",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}

	err = database.DB.WithContext(ctx).
		Delete(&movie).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "could not delete movie",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "movie deleted successfully",
	})
}

func HardDeleteMovie(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	var movie models.Movie

	err := database.DB.WithContext(ctx).
		Unscoped().
		First(&movie, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"message": "movie not found",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}

	err = database.DB.WithContext(ctx).
		Unscoped().
		Delete(&movie).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "could not permanently delete movie",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "movie permanently deleted",
	})
}

func BatchCreateMovies(c *gin.Context) {
	ctx := c.Request.Context()

	var movies []models.Movie

	err := c.ShouldBindJSON(&movies)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid data",
		})
		return
	}

	err = database.DB.WithContext(ctx).
		CreateInBatches(&movies, 10).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "could not create movies",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "movies created successfully",
	})
}

func GetHighRatedMovies(c *gin.Context) {
	ctx := c.Request.Context()

	var movies []models.Movie

	err := database.DB.WithContext(ctx).
		Scopes(models.HighRated).
		Preload("Genres").
		Find(&movies).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}

	c.JSON(200, movies)
}

func GetLatestMovies(c *gin.Context) {
	ctx := c.Request.Context()

	var movies []models.Movie

	err := database.DB.WithContext(ctx).
		Scopes(models.Recentmovies).
		Preload("Genres").
		Find(&movies).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}

	c.JSON(200, movies)
}

// --------------------------------------------------
// REVIEWS
// --------------------------------------------------

func GetReviewsByMovie(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	var reviews []models.Review

	err := database.DB.WithContext(ctx).
		Where("movie_id = ?", id).
		Find(&reviews).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "could not get reviews",
		})
		return
	}

	c.JSON(200, reviews)
}

// --------------------------------------------------
// STATS
// --------------------------------------------------

func GetStats(c *gin.Context) {
	ctx := c.Request.Context()

	var movieCount int64
	var userCount int64
	var reviewCount int64

	err := database.DB.WithContext(ctx).
		Model(&models.Movie{}).
		Count(&movieCount).Error

	if err != nil {
		c.JSON(500, gin.H{"message": "internal error"})
		return
	}

	err = database.DB.WithContext(ctx).
		Model(&models.User{}).
		Count(&userCount).Error

	if err != nil {
		c.JSON(500, gin.H{"message": "internal error"})
		return
	}

	err = database.DB.WithContext(ctx).
		Model(&models.Review{}).
		Count(&reviewCount).Error

	if err != nil {
		c.JSON(500, gin.H{"message": "internal error"})
		return
	}

	c.JSON(200, gin.H{
		"total_movies": movieCount,
		"total_users":  userCount,
		"total_reviews": reviewCount,
	})
}

// --------------------------------------------------
// AUTH
// --------------------------------------------------

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=2"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	ctx := c.Request.Context()

	var input RegisterRequest

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	// Check duplicate email
	var existingUser models.User

	err = database.DB.WithContext(ctx).
		Where("email = ?", input.Email).
		First(&existingUser).Error

	if err == nil {
		c.JSON(409, gin.H{
			"error": "email already registered",
		})
		return
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(500, gin.H{
			"error": "internal error",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "could not hash password",
		})
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	err = database.DB.WithContext(ctx).
		Create(&user).Error

	if err != nil {
		c.JSON(500, gin.H{
			"error": "could not create user",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "user registered successfully",
	})
}

func Login(c *gin.Context) {
	ctx := c.Request.Context()

	var input LoginRequest

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	var user models.User

	err = database.DB.WithContext(ctx).
		Where("email = ?", input.Email).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(401, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": "internal error",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil {
		c.JSON(401, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		c.JSON(500, gin.H{
			"error": "jwt secret not configured",
		})
		return
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(
		[]byte(secret),
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "could not generate token",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}

// --------------------------------------------------
// CURRENT USER
// --------------------------------------------------

func GetCurrentUser(c *gin.Context) {
	ctx := c.Request.Context()

	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	userID, ok := userIDValue.(uint)

	if !ok {
		c.JSON(401, gin.H{
			"message": "invalid user",
		})
		return
	}

	var user models.User

	err := database.DB.WithContext(ctx).
		First(&user, userID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"message": "user not found",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}

	c.JSON(200, gin.H{
		"id":        user.ID,
		"name":      user.Name,
		"email":     user.Email,
		"joined_at": user.CreatedAt,
	})
}

// --------------------------------------------------
// TRACKING
// --------------------------------------------------

type TrackingRequest struct {
	Status     string `json:"status"`
	UserRating int    `json:"userRating"`
	Review     string `json:"review"`
}

func SaveTracking(c *gin.Context) {
	ctx := c.Request.Context()

	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	userID, ok := userIDValue.(uint)

	if !ok {
		c.JSON(401, gin.H{
			"message": "invalid user",
		})
		return
	}

	id := c.Param("id")

	var movie models.Movie

	err := database.DB.WithContext(ctx).
		First(&movie, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"message": "movie not found",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}

	var input TrackingRequest

	err = c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid input",
		})
		return
	}

	if input.Status == "" {
		c.JSON(400, gin.H{
			"message": "status is required",
		})
		return
	}

	if input.Status != "Watched" &&
		input.Status != "Watching" &&
		input.Status != "Want to Watch" &&
		input.Status != "Dropped" {

		c.JSON(400, gin.H{
			"message": "invalid status",
		})
		return
	}

	if input.UserRating < 1 || input.UserRating > 10 {
		c.JSON(400, gin.H{
			"message": "rating must be between 1 and 10",
		})
		return
	}

	// -------------------------
	// STATUS
	// -------------------------

	var tracking models.UserMovie

	err = database.DB.WithContext(ctx).
		Where(
			"user_id = ? AND movie_id = ?",
			userID,
			movie.ID,
		).
		First(&tracking).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		tracking = models.UserMovie{
			UserID:  userID,
			MovieID: movie.ID,
			Status:  input.Status,
		}

		err = database.DB.WithContext(ctx).
			Create(&tracking).Error

		if err != nil {
			c.JSON(500, gin.H{
				"message": "could not save status",
			})
			return
		}
	} else if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	} else {
		tracking.Status = input.Status

		err = database.DB.WithContext(ctx).
			Save(&tracking).Error

		if err != nil {
			c.JSON(500, gin.H{
				"message": "could not update status",
			})
			return
		}
	}

	// -------------------------
	// RATING + REVIEW
	// -------------------------

	var review models.Review

	err = database.DB.WithContext(ctx).
		Where(
			"user_id = ? AND movie_id = ?",
			userID,
			movie.ID,
		).
		First(&review).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		review = models.Review{
			UserID:  userID,
			MovieID: movie.ID,
			Rating:  input.UserRating,
			Comment: input.Review,
		}

		err = database.DB.WithContext(ctx).
			Create(&review).Error

		if err != nil {
			c.JSON(500, gin.H{
				"message": "could not save review",
			})
			return
		}
	} else if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	} else {
		review.Rating = input.UserRating
		review.Comment = input.Review

		err = database.DB.WithContext(ctx).
			Save(&review).Error

		if err != nil {
			c.JSON(500, gin.H{
				"message": "could not update review",
			})
			return
		}
	}

	// -------------------------
	// UPDATE MOVIE AVERAGE RATING
	// -------------------------

	var avgRating float64

	err = database.DB.WithContext(ctx).
		Model(&models.Review{}).
		Where("movie_id = ?", movie.ID).
		Select("AVG(rating)").
		Scan(&avgRating).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "could not calculate average rating",
		})
		return
	}

	err = database.DB.WithContext(ctx).
		Model(&models.Movie{}).
		Where("id = ?", movie.ID).
		Update("rating", int(avgRating)).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "could not update movie rating",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "tracking saved successfully",
	})
}

// --------------------------------------------------
// GET TRACKING FOR ONE MOVIE
// --------------------------------------------------

func GetTracking(c *gin.Context) {
	ctx := c.Request.Context()

	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	userID, ok := userIDValue.(uint)

	if !ok {
		c.JSON(401, gin.H{
			"message": "invalid user",
		})
		return
	}

	id := c.Param("id")

	var movie models.Movie

	err := database.DB.WithContext(ctx).
		First(&movie, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"message": "movie not found",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}

	response := gin.H{
		"status":     "",
		"userRating": nil,
		"review":     "",
	}

	var tracking models.UserMovie

	err = database.DB.WithContext(ctx).
		Where(
			"user_id = ? AND movie_id = ?",
			userID,
			movie.ID,
		).
		First(&tracking).Error

	if err == nil {
		response["status"] = tracking.Status
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(500, gin.H{
			"message": "could not get tracking",
		})
		return
	}

	var review models.Review

	err = database.DB.WithContext(ctx).
		Where(
			"user_id = ? AND movie_id = ?",
			userID,
			movie.ID,
		).
		First(&review).Error

	if err == nil {
		response["userRating"] = review.Rating
		response["review"] = review.Comment
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(500, gin.H{
			"message": "could not get review",
		})
		return
	}

	c.JSON(200, response)
}

// --------------------------------------------------
// MY MOVIES
// --------------------------------------------------

type MyMovieResponse struct {
	ID          uint           `json:"id"`
	Title       string         `json:"title"`
	Rating      int            `json:"rating"`
	ReleaseYear int            `json:"release_year"`
	Description string         `json:"description"`
	Genres      []models.Genre `json:"genres"`
	Status      string         `json:"status"`
	UserRating  int            `json:"userRating"`
	Review      string         `json:"review"`
}

func GetMyMovies(c *gin.Context) {
	ctx := c.Request.Context()

	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	userID, ok := userIDValue.(uint)

	if !ok {
		c.JSON(401, gin.H{
			"message": "invalid user",
		})
		return
	}

	var trackedMovies []models.UserMovie

	err := database.DB.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("updated_at DESC").
		Find(&trackedMovies).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "could not get tracked movies",
		})
		return
	}

	result := []MyMovieResponse{}

	for _, tracking := range trackedMovies {
		var movie models.Movie

		err = database.DB.WithContext(ctx).
			Preload("Genres").
			First(&movie, tracking.MovieID).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			continue
		}

		if err != nil {
			c.JSON(500, gin.H{
				"message": "could not get movie",
			})
			return
		}

		var review models.Review

		err = database.DB.WithContext(ctx).
			Where(
				"user_id = ? AND movie_id = ?",
				userID,
				movie.ID,
			).
			First(&review).Error

		if err != nil &&
			!errors.Is(err, gorm.ErrRecordNotFound) {

			c.JSON(500, gin.H{
				"message": "could not get review",
			})
			return
		}

		item := MyMovieResponse{
			ID:          movie.ID,
			Title:       movie.Title,
			Rating:      movie.Rating,
			ReleaseYear: movie.Release_year,
			Description: movie.Description,
			Genres:      movie.Genres,
			Status:      tracking.Status,
			UserRating:  review.Rating,
			Review:      review.Comment,
		}

		result = append(result, item)
	}

	c.JSON(200, result)
}