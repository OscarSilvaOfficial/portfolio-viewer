package domain

func FactoryViewer(userAgent string) Viewer {
	viewer := Viewer{}
	viewer.SetUserAgent(userAgent)
	return viewer
}
