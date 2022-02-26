package domain

type Viewer struct {
	userAgent string
}

func (self *Viewer) SetUserAgent(userAgent string) *Viewer {
	self.userAgent = userAgent
	return self
}

func (self *Viewer) UserAgent() string {
	return self.userAgent
}

func FactoryViewer(userAgent string) Viewer {
	viewer := Viewer{}
	viewer.SetUserAgent(userAgent)
	return viewer
}
