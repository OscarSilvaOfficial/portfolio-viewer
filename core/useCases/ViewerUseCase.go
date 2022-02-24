package usecases

import "github.com/OscarSilvaOfficial/portfolio-viewer/core/domain"

func AddNewViewer(userAgent string) map[string]string {
	viewer := domain.FactoryViewer(userAgent)
	return map[string]string{"viewer": viewer.UserAgent()}
}
