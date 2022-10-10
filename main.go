package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alpocketmonster/GoHttpServer/controller"
	"github.com/gin-gonic/gin"
)

type SigHandler interface {
	SighupHandler()
}

func setupRouter() (*gin.Engine, SigHandler) {
	r := gin.Default()
	controller := controller.NewController()

	r.GET("/auth", func(ctx *gin.Context) {
		//Если content type = ct ->200 || 400
		//или Request.Header.Get("Content-Type")

		code, err := controller.Validate(ctx.ContentType(), ctx.GetHeader("X-Original-Uri"))

		switch code {
		case http.StatusOK: //200
			ctx.JSON(http.StatusOK, gin.H{"message": "Request is valid!"})
		case http.StatusBadRequest: //400
			ctx.JSON(http.StatusBadRequest, gin.H{"message": string(err.Error())})
		}
	})
	return r, controller
}

func main() {
	var wg sync.WaitGroup
	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

	r, controller := setupRouter()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	quit := make(chan os.Signal)

	wg.Add(1)
	go func() {
		defer wg.Done()
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen failed: %s\n", err)
			close(quit)
		}
		//time.Sleep(3 * time.Second)
		log.Println("Server done")

	}()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	time.Sleep(3 * time.Second)
	// 	log.Println("Server hello")
	// }()
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for {
		s, ok := <-quit
		if !ok {
			break
		}
		switch s {
		case syscall.SIGHUP:
			controller.SighupHandler()
		case syscall.SIGINT, syscall.SIGTERM:
			log.Println("term")
			log.Println("Shutdown Server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}
			log.Println("Waiting")
			wg.Wait()
			log.Println("Server exiting")
			return
		}
	}
}
