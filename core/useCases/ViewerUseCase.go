package usecases

import (
	"github.com/OscarSilvaOfficial/portfolio-viewer/core/domain"
	"github.com/OscarSilvaOfficial/portfolio-viewer/ports"
)

type RMap struct{
	value map[string]string
}

func AddNewViewer(userAgent string, repository ports.ViewerRepository) map[string]string {
	viewer := domain.FactoryViewer(userAgent)
	repository.AddNewViewer(viewer)
	return Map("viewer", viewer.UserAgent())
}

func Map(key string, value string) map[string]string {
	return map[string]string{key: value}
}