package database

import (
	"cinetrack/models"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
func Connect(){
	dsn:="host=localhost password=saatwik2404 dbname=moviedb user=postgres sslmode=disable port=5432"

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