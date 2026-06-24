package main

import (
	"cinetrack/database"
	"github.com/gin-gonic/gin"
)


func main(){

database.Connect()
router:=gin.Default()
router.Run(":8080")

}