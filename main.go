package main

import (
    "cinetrack/database"
    "cinetrack/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    database.Connect()
    router := gin.Default()
    routes.SetupRoutes(router)
    router.Run(":9092")
}