package adapters

import (
	usecases "github.com/OscarSilvaOfficial/portfolio-viewer/core/useCases"
	"github.com/OscarSilvaOfficial/portfolio-viewer/ports"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	repository ports.ViewerRepository
}

func (self *Controller) AddNewViewer(g *gin.Context) {
	userAgent := g.Request.UserAgent()
	viewer := usecases.AddNewViewer(userAgent, self.repository)
	self.Response(201, viewer, g)
}

func (self *Controller) Response(statusCode int, data map[string]string, g *gin.Context) {
	response := gin.H{"data": data}
	g.JSON(statusCode, response)
}

func (self *Controller) SetRepository(repository ports.ViewerRepository) *Controller {
	self.repository = repository
	return self
}

func FactoryController(repository ports.ViewerRepository) Controller {
	controller := Controller{}
	controller.SetRepository(repository)
	return controller
}
