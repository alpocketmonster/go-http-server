package controller

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type controller struct {
	sighup chan os.Signal
	ct     string
}

func NewController() *controller {
	var cntr controller
	cntr.sighup = make(chan os.Signal)
	go cntr.sighupHandler()

	return &cntr
}

func (p *controller) sighupHandler() {
	for {
		signal, ok := <-p.sighup
		if ok {
			log.Println("sighup", signal.String())
		}
	}
}

func (p *controller) GetSighupChan() chan os.Signal {
	return p.sighup
}
func (p *controller) SetCt(newCt string) {
	p.ct = newCt
}

func (c *controller) AuthMiddleware(ctx *gin.Context) {
	//Если content type = ct ->200 || 400
	//или Request.Header.Get("Content-Type")
	if ctx.ContentType() == c.ct {
		ctx.JSON(http.StatusOK, gin.H{"message": "Content Type is valid!"})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Content Type is invalid!"})
}
