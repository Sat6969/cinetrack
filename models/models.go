package models

import (
	"cinetrack/database"
	"cinetrack/models"

	"gorm.io/gorm"
)

type Movie struct{
	gorm.Model
	Title string `json:"title" gorm:"not null;uniqueIndex"`
	Rating int	`json:"rating" gorm:"not null;check:rating>=0 AND rating<=10"`
	Release_year int `json:"release_year" gorm:"not null"`
	Description string	`json:"description" gorm:"not null"`
	Reviews []Review `json:"review"`
	Genres []Genre `json:"genres" gorm:"many2many:movie_genres"`
}

type User struct{
	gorm.Model
	Name string	`json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null"`
	Reviews []Review  `json:"review"`
}

type Genre struct{
	gorm.Model 
	Name string `json:"name" gorm:"not null;unique"`
	Movies []Movie `json:"movies" gorm:"many2many:movie_genres"`
}

type Review struct{
	gorm.Model
	UserID int 	`json:"user_id" gorm:"not null"`
	MovieID int `json:"movie_id" gorm:"not null"`
	Rating int `json:"rating" gorm:"not null"`
	Comment string `json:"comment"`
}

func HighRated(db *gorm.DB) (*gorm.DB){

	return db.Where("rating > ?",7)
}

func Recentmovies(db *gorm.DB) (*gorm.DB){
	return db.Where("release_year > ?",2020)
}