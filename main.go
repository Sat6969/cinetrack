package main

import (
	"cinetrack/database"
	"cinetrack/routes"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
    database.Connect()
    router := gin.Default()
    routes.SetupRoutes(router)

    srv:=&http.Server{
        Addr:"9092",
        Handler:router,
    }
    go func(){
         err:=srv.ListenAndServe()
         if (err!=nil && err!=http.ErrServerClosed){
            panic(err)
         }
    }()
    quit:=make(chan os.Signal,1)
    signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
    <-quit

    ctx,canel:=context.WithTimeout(context.Background(),5*time.Second)
    defer canel()
    srv.Shutdown(ctx)
    router.Run(":9092")
}