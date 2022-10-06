package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alpocketmonster/GoHttpServer/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	return gin.New()
}

func TestHomepageHandler(t *testing.T) {
	controller := new(controller.Controller)
	router := SetUpRouter()
	router.GET("/auth", controller.AuthMiddleware)
	request, _ := http.NewRequest("GET", "/auth", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	//responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHomepageHandlerWrongCt(t *testing.T) {
	controller := new(controller.Controller)
	controller.SetCt("wrong")

	router := SetUpRouter()
	router.GET("/auth", controller.AuthMiddleware)
	request, _ := http.NewRequest("GET", "/auth", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHomepageHandlerWrongContentType(t *testing.T) {
	controller := new(controller.Controller)
	router := SetUpRouter()
	router.GET("/auth", controller.AuthMiddleware)

	request, _ := http.NewRequest("GET", "/auth", nil)
	request.Header.Set("Content-Type", "wrong")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func BenchmarkTestHomepageHandler(b *testing.B) {
	controller := new(controller.Controller)
	router := SetUpRouter()
	router.GET("/auth", controller.AuthMiddleware)
	request, _ := http.NewRequest("GET", "/auth", nil)
	w := httptest.NewRecorder()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		router.ServeHTTP(w, request)
	}
}
