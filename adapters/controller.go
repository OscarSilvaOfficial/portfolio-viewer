package adapters

import (
	usecases "github.com/OscarSilvaOfficial/portfolio-viewer/core/useCases"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (controller *Controller) AddNewViewer(g *gin.Context) {
	userAgent := g.Request.UserAgent()
	viewer := usecases.AddNewViewer(userAgent)
	controller.Response(201, viewer, g)
}

func (controller *Controller) Response(statusCode int, data map[string]string, g *gin.Context) {
	response := gin.H{"data": data}
	g.JSON(statusCode, response)
}

func FactoryController() Controller {
	controller := Controller{}
	return controller
}
