package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	ct string
}

func (p *Controller) SetCt(newCt string) {
	p.ct = newCt
}

func (c *Controller) AuthMiddleware(ctx *gin.Context) {
	//Если content type = ct ->200 || 400
	//или Request.Header.Get("Content-Type")
	if ctx.ContentType() == c.ct {
		ctx.JSON(http.StatusOK, gin.H{"message": "Content Type is valid!"})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Content Type is invalid!"})
}
