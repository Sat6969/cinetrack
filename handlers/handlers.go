package handlers

import (
	"cinetrack/database"
	"cinetrack/models"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func handler1(c *gin.Context){
	var movies []models.Movie
	ctx:=c.Request.Context()
	err:=database.DB.WithContext(ctx).Raw("select * from movies").Scan(&movies).Error

	if(errors.Is(err,gorm.ErrRecordNotFound)){
		c.JSON(404,gin.H{
			"message":"data not found",
		})
		return
	}
	if(err!=nil){
		c.JSON(500,gin.H{
			"message":"internal error",
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"data found",
		"data":movies,
	})
	
}

func handler2(c *gin.Context){
	id:=c.Param("id");
	ctx:=c.Request.Context()
	var movies models.Movie
	err:=database.DB.WithContext(ctx).Raw("select * from movies where id=?",id).Scan(&movies).Error


	if(errors.Is(err,gorm.ErrRecordNotFound)){
		c.JSON(404,gin.H{
			"message":"invalid data",
		})
		return
	}

	if(err!=nil){
		c.JSON(500,gin.H{
			"message":"internal issue",
		})
		return
	}
	
	var genres []models.Genre

	err2:=database.DB.WithContext(ctx).Raw("select genres.* from movie_genres join genres on genres.id=movie_genres.genre_id where movie_id=?",id).Scan(&genres).Error

	if(err2!=nil){
		c.JSON(500,gin.H{
			"message":"internal error",
		})
		return
	}

	var reviews []models.Review

	err3:=database.DB.WithContext(ctx).Raw("select reviews.* from reviews join movies on reviews.movie_id=movies.id where movies.id=?",id).Scan(&reviews).Error

	if(err3!=nil){
		c.JSON(500,gin.H{
			"message":"internal error",
		})
		return
	}

	movies.Genres=genres
	movies.Reviews=reviews

	c.JSON(200,gin.H{
		"message":"data found",
		"movie":movies,
	})
	

}

func handler3(c *gin.Context){

	    var input struct {
        Title       string   `json:"title"`
        Rating      int      `json:"rating"`
        ReleaseYear int      `json:"release_year"`
        Description string   `json:"description"`
        Genres      []string `json:"genres"`
    }

	 err:=c.ShouldBindJSON(&input)

	if(err!=nil){
		c.JSON(400,gin.H{
			"message":"invalid data",
		})
		return
	}
	
	ctx:=c.Request.Context()

	movie:=models.Movie{
		Title: input.Title,
		Rating: input.Rating,
		Description: input.Description,
	}

	for _,name:=range input.Genres{
		var genres models.Genre
		database.DB.WithContext(ctx).Where("name=?",name).FirstOrCreate(&genres)
		movie.Genres = append(movie.Genres, genres)
	}

	err3:=database.DB.WithContext(ctx).Create(&movie).Error

	 if(err3!=nil){
		c.JSON(500,gin.H{
			"message":"internal error",
		})
		return
	 }

	 c.JSON(200,gin.H{
		"message":"sucesfully created",
	 })


}

func handler4(c *gin.Context){

	id:=c.Param("id");
	ctx:=c.Request.Context()
	var movie models.Movie

	err2:=database.DB.WithContext(ctx).First(&movie,id).Error
	if(err2!=nil){
		c.JSON(500,gin.H{
			"message":"internal error",
		})
		return
	}

	var updated_movie models.Movie

	err:=c.ShouldBindJSON(&updated_movie)

	if(err!=nil){
		c.JSON(404,gin.H{
			"message":"invlid input",
		})
		return
	}
	err3:=database.DB.WithContext(ctx).Model(&models.Movie{}).Where("id=?",id).Updates(&updated_movie).Error

	if(err3!=nil){
		c.JSON(500,gin.H{
			"message":"invalid input",
		})
		return
	}

	c.JSON(200,gin.H{
		"message":"updated sucesfully",
	})



}

func handler5(c *gin.Context){
	id:=c.Param("id");
	ctx:=c.Request.Context()
	var movies models.Movie

	err:=database.DB.WithContext(ctx).First(&movies,id).Error

	if(errors.Is(err,gorm.ErrRecordNotFound)){
		c.JSON(404,gin.H{
			"message":"invalid data",
		})
		return
	}

	if(err!=nil){
		c.JSON(500,gin.H{
			"message":"internal error",
		})
		return
	}

	err2:=database.DB.WithContext(ctx).Delete(&movies,id).Error

	if(err2!=nil){
		c.JSON(500,gin.H{
			"message":"internal error",
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"sucesfully deleted",
	})

}

func handler6(c *gin.Context){

	var users []models.User
	ctx:=c.Request.Context()
	err:=database.DB.WithContext(ctx).Model(&models.User{}).Find(&users).Error

	if(err!=nil){
		c.JSON(500,gin.H{
			"message":"internal error",
		})
		return
	}

	c.JSON(200,gin.H{
		"message":"sucessfull",
		"users":users,
	})

}

func handler7(c *gin.Context){
	ctx:=c.Request.Context()

	id:=c.Param("id")

	var user models.User

	err:=database.DB.WithContext(ctx).Preload("Reviews").First(&user,id).Error

	if(errors.Is(err,gorm.ErrRecordNotFound)){
		c.JSON(404,gin.H{
			"message":"record not found",
		})
	}
	if(err!=nil){
		c.JSON(500,gin.H{
			"message":"internal error",
		})
		return
	}
	
	c.JSON(200,gin.H{
		"message":user,
	})

}

func hanlder8(c *gin.Context){
	ctx:=c.Request.Context()
	var input struct{
		Name string	`json:"name"`
		Email string	`json:"email"`
	}

	err:=c.ShouldBindJSON(&input)

	user:=models.User{
		Name: input.Name,
		Email: input.Email,
	}

	if(err!=nil){
		c.JSON(404,gin.H{
			"messsage":"invalid input",
		})
		return
	}

	err2:=database.DB.WithContext(ctx).Create(&user).Error
	if(err2!=nil){
		c.JSON(500,gin.H{
			"message":"internal error",
		})
		return
	}

	c.JSON(200,gin.H{
		"message":"sucesfully created",
	})

}

func handler8(c *gin.Context){
	id:=c.Param("id")

	var user models.User
	
	ctx:=c.Request.Context()

	err:=database.DB.WithContext(ctx).First(&user,id).Error

	if(errors.Is(err,gorm.ErrRecordNotFound)){
		c.JSON(404,gin.H{
			"message":"invalid data",
		})
		return
	}

	if(err!=nil){
		c.JSON(500,gin.H{
			"message":"something went wrong",
		})
		return
	}

	err2:=database.DB.WithContext(ctx).Model(&models.User{}).Delete(&user,id).Error

	if(err2!=nil){
		c.JSON(500,gin.H{
			"message":"something went wrong",
		})
		return
	}

	c.JSON(200,gin.H{
		"message":"succesfully deleted",
	})

	
}

func handler9(c *gin.Context){
	ctx:=c.Request.Context()

	var review models.Review

	err:=c.ShouldBindJSON(&review)

	
	if(err!=nil){
		c.JSON(404,gin.H{
			"message":"invalid data",
		})
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
		c.JSON(500,gin.H{
			"message":"something went wrong",
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"succesfull",
	})
	
}

func handler10(c *gin.Context){
	id:=c.Param("id")

	var reviews []models.Review

	ctx:=c.Request.Context()

	err:=database.DB.WithContext(ctx).Where("movie_id=?",id).Scan(&reviews).Error

	if(err!=nil){
		c.JSON(500,gin.H{
			"message":"something went wrong",
		})
		return
	}

	c.JSON(200,gin.H{
		"message":reviews,
	})


}