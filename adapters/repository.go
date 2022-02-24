package adapters

import "github.com/OscarSilvaOfficial/portfolio-viewer/ports"

type Repository struct {
	driver ports.ViewerRepository
}

func (r *Repository) AddNewViewer(userAgent string) map[string]string {
	return r.driver.AddNewViewer(userAgent)
}
