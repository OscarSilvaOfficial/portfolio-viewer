package adapters

import (
	"github.com/OscarSilvaOfficial/portfolio-viewer/core/domain"
	"github.com/OscarSilvaOfficial/portfolio-viewer/ports"
)

type Repository struct {
	driver ports.NoSQLPort
}

func (r *Repository) AddNewViewer(viewer domain.Viewer) {
	document := map[string]string{"viewer": viewer.UserAgent()}
	r.driver.CreateDocument(document)
}

func (r *Repository) SetDriver(driver ports.NoSQLPort) Repository {
	r.driver = driver
	return *r
}

func FactoryRepository(driver ports.NoSQLPort) Repository {
	repository := Repository{}
	repository.SetDriver(driver)
	return repository
}
