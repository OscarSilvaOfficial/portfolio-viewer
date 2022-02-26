package main

import (
	"github.com/OscarSilvaOfficial/portfolio-viewer/adapters"
	"github.com/OscarSilvaOfficial/portfolio-viewer/infra/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := db.FactoryMongo("mongodb://127.0.0.1:27017", "portfolio", "viewers")
	repository := adapters.FactoryRepository(&db)
	controller := adapters.FactoryController(&repository)
	r.GET("/", controller.AddNewViewer)
	r.Run()
}
