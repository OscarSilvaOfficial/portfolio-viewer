package usecases

import (
	"github.com/OscarSilvaOfficial/portfolio-viewer/core/domain"
	"github.com/OscarSilvaOfficial/portfolio-viewer/ports"
)

func AddNewViewer(userAgent string, repository ports.ViewerRepository) map[string]string {
	viewer := domain.FactoryViewer(userAgent)
	repository.AddNewViewer(viewer)
	return map[string]string{"viewer": viewer.UserAgent()}
}
