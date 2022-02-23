package adapters

import (
	"github.com/OscarSilvaOfficial/portfolio-viewer/core/factory"
	"github.com/gin-gonic/gin"
)

type Controller struct {}

func (c *Controller) AddNewViewer(g *gin.Context) {
	g.JSON(200, gin.H{"viewers": factory.FactoryViewer("userAgent")})
}