package factory

import "github.com/OscarSilvaOfficial/portfolio-viewer/core/domain"

func FactoryViewer(userAgent string) domain.Viewer {
	viewer := domain.Viewer{}
	viewer.SetUserAgent(userAgent)
	return viewer
}