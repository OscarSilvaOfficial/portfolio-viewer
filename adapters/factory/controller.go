package factory

import "github.com/gin-gonic/gin"
import "github.com/OscarSilvaOfficial/portfolio-viewer/adapters"

func FactoryController(api gin.Context) adapters.Controller {
	controller := adapters.Controller{}
	return controller
}
