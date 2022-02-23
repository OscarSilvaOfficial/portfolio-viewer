package main

import (
	"github.com/OscarSilvaOfficial/portfolio-viewer/adapters/factory"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller := factory.FactoryController()
	r.GET("/", controller.AddNewViewer)
	r.Run()
}
