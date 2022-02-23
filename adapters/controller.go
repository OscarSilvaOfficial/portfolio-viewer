package adapters

import (
	"github.com/OscarSilvaOfficial/portfolio-viewer/core/factory"
	"github.com/gin-gonic/gin"
)

type Controller struct {}

func (controller *Controller) AddNewViewer(g *gin.Context) {
	userAgent := g.Request.UserAgent()
	viewer := factory.FactoryViewer(userAgent)
	response := map[string]string{"viewer": viewer.UserAgent()}
	controller.Response(200, response, g)
}

func (controller *Controller) Response(statusCode int, data map[string]string, g *gin.Context) {
	response := gin.H{"data": data}
	g.JSON(statusCode, response)
}