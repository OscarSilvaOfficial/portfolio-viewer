package ports

import "github.com/OscarSilvaOfficial/portfolio-viewer/core/domain"

type ViewerRepository interface {
	AddNewViewer(viewer domain.Viewer)
}