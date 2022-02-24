package main

import (
	"github.com/OscarSilvaOfficial/portfolio-viewer/adapters"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller := adapters.FactoryController()
	r.GET("/", controller.AddNewViewer)
	r.Run()
}
