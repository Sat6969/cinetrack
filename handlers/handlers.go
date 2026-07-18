package handlers

import (
	"cinetrack/database"
	"cinetrack/models"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func GetAllMovies(c *gin.Context){
	var movies []models.Movie
	ctx:=c.Request.Context()
	err:=database.DB.WithContext(ctx).Preload("Genres").Find(&movies).Error

	if(errors.Is(err,gorm.ErrRecordNotFound)){
		c.JSON(404,gin.H{"message":"data not found"})
		return
	}
	if(err!=nil){
		c.JSON(500,gin.H{"message":"internal error"})
		return
	}
	c.JSON(200,gin.H{"message":"data found","data":movies})
}

func GetMovieByID(c *gin.Context){
	id:=c.Param("id")
	ctx:=c.Request.Context()
	var movies models.Movie
	err:=database.DB.WithContext(ctx).Raw("select * from movies where id=?",id).Scan(&movies).Error

	if(errors.Is(err,gorm.ErrRecordNotFound)){
		c.JSON(404,gin.H{"message":"invalid data"})
		return
	}
	if(err!=nil){
		c.JSON(500,gin.H{"message":"internal issue"})
		return
	}
	
	var genres []models.Genre
	err2:=database.DB.WithContext(ctx).Raw("select genres.* from movie_genres join genres on genres.id=movie_genres.genre_id where movie_id=?",id).Scan(&genres).Error
	if(err2!=nil){
		c.JSON(500,gin.H{"message":"internal error"})
		return
	}

	var reviews []models.Review
	err3:=database.DB.WithContext(ctx).Raw("select reviews.* from reviews join movies on reviews.movie_id=movies.id where movies.id=?",id).Scan(&reviews).Error
	if(err3!=nil){
		c.JSON(500,gin.H{"message":"internal error"})
		return
	}

	movies.Genres=genres
	movies.Reviews=reviews
	c.JSON(200,gin.H{"message":"data found","movie":movies})
}

func CreateMovie(c *gin.Context){
	var input struct {
		Title       string   `json:"title"`
		Rating      int      `json:"rating"`
		ReleaseYear int      `json:"release_year"`
		Description string   `json:"description"`
		Genres      []string `json:"genres"`
	}

	err:=c.ShouldBindJSON(&input)
	if(err!=nil){
		c.JSON(400,gin.H{"message":"invalid data"})
		return
	}
	fmt.Println(input.Genres)
	
	ctx:=c.Request.Context()
	movie:=models.Movie{
		Title: input.Title,
		Rating: input.Rating,
		Description: input.Description,
		Release_year: input.ReleaseYear,
	}
for _, name := range input.Genres {
    var genres models.Genre
    result := database.DB.WithContext(ctx).Where("name = ?", name).FirstOrCreate(&genres,models.Genre{Name:name})
    if result.Error != nil {
        c.JSON(500, gin.H{"message": "failed to create genre"})
        return
    }
    movie.Genres = append(movie.Genres, genres)
}

	err3:=database.DB.WithContext(ctx).Create(&movie).Error
	if(err3!=nil){
		c.JSON(500,gin.H{"message":"internal error"})
		return
	}
	c.JSON(201,gin.H{"message":"sucesfully created"})
}

func UpdateMovie(c *gin.Context){
	id:=c.Param("id")
	ctx:=c.Request.Context()
	var movie models.Movie

	err2:=database.DB.WithContext(ctx).First(&movie,id).Error
	if(err2!=nil){
		c.JSON(500,gin.H{"message":"internal error"})
		return
	}

	var updated_movie models.Movie
	err:=c.ShouldBindJSON(&updated_movie)
	if(err!=nil){
		c.JSON(404,gin.H{"message":"invlid input"})
		return
	}

	err3:=database.DB.WithContext(ctx).Model(&models.Movie{}).Where("id=?",id).Updates(&updated_movie).Error
	if(err3!=nil){
		c.JSON(500,gin.H{"message":"invalid input"})
		return
	}
	c.JSON(200,gin.H{"message":"updated sucesfully"})
}

func DeleteMovie(c *gin.Context){
	id:=c.Param("id")
	ctx:=c.Request.Context()
	var movies models.Movie

	err:=database.DB.WithContext(ctx).First(&movies,id).Error
	if(errors.Is(err,gorm.ErrRecordNotFound)){
		c.JSON(404,gin.H{"message":"invalid data"})
		return
	}
	if(err!=nil){
		c.JSON(500,gin.H{"message":"internal error"})
		return
	}

	err2:=database.DB.WithContext(ctx).Delete(&movies,id).Error
	if(err2!=nil){
		c.JSON(500,gin.H{"message":"internal error"})
		return
	}
	c.JSON(200,gin.H{"message":"sucesfully deleted"})
}

func GetAllUsers(c *gin.Context){
	var users []models.User
	ctx:=c.Request.Context()
	err:=database.DB.WithContext(ctx).Model(&models.User{}).Find(&users).Error
	if(err!=nil){
		c.JSON(500,gin.H{"message":"internal error"})
		return
	}
	c.JSON(200,gin.H{"message":"sucessfull","users":users})
}

func GetUserByID(c *gin.Context){
	ctx:=c.Request.Context()
	id:=c.Param("id")
	var user models.User

	err:=database.DB.WithContext(ctx).Preload("Reviews").First(&user,id).Error
	if(errors.Is(err,gorm.ErrRecordNotFound)){
		c.JSON(404,gin.H{"message":"record not found"})
		return
	}
	if(err!=nil){
		c.JSON(500,gin.H{"message":"internal error"})
		return
	}
	c.JSON(200,gin.H{"message":user})
}

func CreateUser(c *gin.Context){
	ctx:=c.Request.Context()
	var input struct{
		Name string  `json:"name"`
		Email string `json:"email"`
	}

	err:=c.ShouldBindJSON(&input)
	if(err!=nil){
		c.JSON(404,gin.H{"messsage":"invalid input"})
		return
	}
	
	user:=models.User{
		Name: input.Name,
		Email: input.Email,
	}

	err2:=database.DB.WithContext(ctx).Create(&user).Error
	if(err2!=nil){
		c.JSON(500,gin.H{"message":"internal error"})
		return
	}
	c.JSON(200,gin.H{"message":"sucesfully created"})
}

func DeleteUser(c *gin.Context){
	id:=c.Param("id")
	ctx:=c.Request.Context()
	var user models.User

	err:=database.DB.WithContext(ctx).First(&user,id).Error
	if(errors.Is(err,gorm.ErrRecordNotFound)){
		c.JSON(404,gin.H{"message":"invalid data"})
		return
	}
	if(err!=nil){
		c.JSON(500,gin.H{"message":"something went wrong"})
		return
	}

	err2:=database.DB.WithContext(ctx).Model(&models.User{}).Delete(&user,id).Error
	if(err2!=nil){
		c.JSON(500,gin.H{"message":"something went wrong"})
		return
	}
	c.JSON(200,gin.H{"message":"succesfully deleted"})
}

func CreateReview(c *gin.Context){
	ctx:=c.Request.Context()
	var review models.Review

	err:=c.ShouldBindJSON(&review)
	if(err!=nil){
		c.JSON(404,gin.H{"message":"invalid data"})
		return
	}

	err2:=database.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err:=tx.Create(&review).Error
		if(err!=nil){
			return err
		} 
		var avgrating float64
		tx.Model(&models.Review{}).Where("movie_id=?",review.MovieID).Select("AVG(rating)").Scan(&avgrating)

		if err:=tx.Model(&models.Movie{}).Where("id=?",review.MovieID).Update("rating",avgrating).Error; err!=nil{
			return err
		}
		return nil
	})

	if(err2!=nil){
		c.JSON(500,gin.H{"message":"something went wrong"})
		return
	}
	c.JSON(200,gin.H{"message":"succesfull"})
}

func GetReviewsByMovie(c *gin.Context){
	id:=c.Param("id")
	ctx:=c.Request.Context()
	var reviews []models.Review

	err:=database.DB.WithContext(ctx).Where("movie_id=?",id).Find(&reviews).Error
	if(err!=nil){
		c.JSON(500,gin.H{"message":"something went wrong"})
		return
	}
	c.JSON(200,gin.H{"message":reviews})
}

func GetHighRatedMovies(c *gin.Context){
	ctx:=c.Request.Context()

	var movies []models.Movie

	err:=database.DB.WithContext(ctx).Scopes(models.HighRated).Preload("Genres").Preload("Reviews").Find(&movies).Error

	if(err!=nil){
		c.JSON(500,gin.H{
			"message":"internal error",
		})
		return
	}

	c.JSON(200,gin.H{
		"message":movies,
	})
}

func GetLatestMovies(c *gin.Context){
	ctx:=c.Request.Context()

	var movie []models.Movie

	err:=database.DB.WithContext(ctx).Scopes(models.Recentmovies).Preload("Genres").Preload("Reviews").Find(&movie).Error

	if(err!=nil){
		c.JSON(500,gin.H{
			"message":"internal error",
		})
		return
	}

	c.JSON(200,gin.H{
		"message":movie,
	})
}

func BatchCreateMovies(c *gin.Context){
	ctx:=c.Request.Context()
	var movies []models.Movie

	err:=c.ShouldBindJSON(&movies)
	if(err!=nil){
		c.JSON(404,gin.H{
			"message":"invalid data",
		})
		return
	}

	err2:=database.DB.WithContext(ctx).CreateInBatches(&movies,10).Error

	if(err2!=nil){
		c.JSON(500,gin.H{
			"message":"something went wrong",
		})
		return
	}

	c.JSON(200,gin.H{
		"message":"sucessfully created",
	})
}