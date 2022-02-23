package adapters

import (
	"github.com/OscarSilvaOfficial/portfolio-viewer/core/factory"
	"github.com/gin-gonic/gin"
)

type Controller struct {}

func (controller *Controller) AddNewViewer(g *gin.Context) {
	userAgent := g.Request.UserAgent()
	viewer := factory.FactoryViewer(userAgent)
	response := gin.H{"viewer": viewer.UserAgent()}
	g.JSON(200, response)
}
