package database

import (
	"cinetrack/models"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
func Connect(){
	godotenv.Load()
	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",os.Getenv("DB_HOST"),os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_NAME"),os.Getenv("DB_PORT"))

	db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if(err!=nil){
		panic(err);
	}
	db.AutoMigrate(&models.User{}, &models.Genre{}, &models.Movie{}, &models.Review{})

	DB=db;

	sqlDB,err2:=db.DB()
	
	if(err2!=nil){
		panic(err2)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)
	

}