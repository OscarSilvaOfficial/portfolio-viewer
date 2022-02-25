package adapters

import "github.com/OscarSilvaOfficial/portfolio-viewer/ports"

type Repository struct {
	driver ports.NoSQLPort
}

func (r *Repository) AddNewViewer(userAgent string) map[string]string {
	document := map[string]string{"viewer": userAgent}
	return r.driver.CreateDocument(document)
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
