package ports

type ViewerRepository interface {
	AddNewViewer(userAgent string) map[string]string
}