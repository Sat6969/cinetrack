package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)


func TestGetAllMovies(t *testing.T){
	router:=gin.Default()
	router.GET("/movies",GetAllMovies)

	w:=httptest.NewRecorder()
	req,_:=http.NewRequest("GET","/movies",nil)
	router.ServeHTTP(w,req)

	if w.Code!=200{
		t.Errorf("expected 200 got %d",w.Code)
	}
}

func TestGetMovieByID(t *testing.T){
	router:=gin.Default()
	router.GET("/movies/:id",GetMovieByID)

	w:=httptest.NewRecorder()

	req,_:=http.NewRequest("GET","/movies/1",nil)

	router.ServeHTTP(w,req)
	if w.Code!=200{
		t.Errorf("expected 200 but got %d",w.Code)
	}

}

func TestCreateMovie(t *testing.T){
	router:=gin.Default()
	router.POST("/movies",CreateMovie)
	body:=[]byte(`{"title":"Test Movie","rating":8,"release_year":2020,"description":"test","genres":["action"]}`)
	w:=httptest.NewRecorder()
	req,_:=http.NewRequest("POST","/movies",bytes.NewBuffer(body))

	req.Header.Set("Content-Type","application/json")
	router.ServeHTTP(w,req)

	if w.Code!=201{
		t.Errorf("expected 201 but got %d",w.Code)
	}
}

func TestCreateReview(t *testing.T){
	router:=gin.Default()
	router.POST("/reviews",CreateReview)

	 body := []byte(`{"user_id":1,"movie_id":1,"rating":9,"comment":"Great movie!"}`)
    
	w:=httptest.NewRecorder()
	req,_:=http.NewRequest("POST","/reviews",bytes.NewBuffer(body))
	router.ServeHTTP(w,req)

	if w.Code!=201{
		t.Errorf("expected 201,got %d",w.Code)
	}

}