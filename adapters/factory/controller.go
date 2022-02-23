package factory

import "github.com/OscarSilvaOfficial/portfolio-viewer/adapters"

func FactoryController() adapters.Controller {
	controller := adapters.Controller{}
	return controller
}
