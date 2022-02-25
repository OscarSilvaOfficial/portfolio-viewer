package adapters

import (
	usecases "github.com/OscarSilvaOfficial/portfolio-viewer/core/useCases"
	"github.com/OscarSilvaOfficial/portfolio-viewer/ports"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	repository ports.ViewerRepository
}

func (controller *Controller) AddNewViewer(g *gin.Context) {
	userAgent := g.Request.UserAgent()
	viewer := usecases.AddNewViewer(userAgent)
	controller.Response(201, viewer, g)
}

func (controller *Controller) Response(statusCode int, data map[string]string, g *gin.Context) {
	response := gin.H{"data": data}
	g.JSON(statusCode, response)
}

func (controller *Controller) SetRepository(repository ports.ViewerRepository) Controller {
	controller.repository = repository
	return *controller
}

func FactoryController(repository ports.ViewerRepository) Controller {
	controller := Controller{}
	controller.SetRepository(repository)
	return controller
}
