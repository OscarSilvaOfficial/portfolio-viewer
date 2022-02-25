package main

import (
	"github.com/OscarSilvaOfficial/portfolio-viewer/adapters"
	"github.com/OscarSilvaOfficial/portfolio-viewer/infra/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := db.FactoryMongo("sdasdad")
	repository := adapters.FactoryRepository(&db)
	controller := adapters.FactoryController(&repository)
	r.GET("/", controller.AddNewViewer)
	r.Run()
}
