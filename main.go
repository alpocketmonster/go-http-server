package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alpocketmonster/GoHttpServer/controller"
	"github.com/alpocketmonster/GoHttpServer/tools"
	"github.com/gin-gonic/gin"
)

type SigHandler interface {
	SighupHandler()
	String() string
}

func setupRouter() (*gin.Engine, SigHandler) {
	r := gin.Default()
	controller := controller.NewController()

	r.GET("/auth", func(ctx *gin.Context) {
		//Если content type = ct ->200 || 400

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
	logger := tools.NewLogExtended()
	logger.SetLogLevel(tools.LogLevelInfo)
	logger.Warnln("DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

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
			logger.Errorln(fmt.Sprintf("listen failed: %s\n", err))
			close(quit)
		}
		//time.Sleep(3 * time.Second)
		logger.Infoln("Server done")
		logger.Infoln(controller.String())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		logger.Infoln("Server hello")
		logger.Infoln(controller.String())
	}()

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
			logger.Infoln("term")
			logger.Infoln("Shutdown Server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				logger.Errorln(fmt.Sprint("Server Shutdown:", err))
			}
			logger.Infoln("Waiting")
			wg.Wait()
			logger.Infoln("Server exiting")
			return
		}
	}
}
