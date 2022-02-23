package domain

type Viewer struct {
	userAgent string
}

func (v *Viewer) SetUserAgent(userAgent string) Viewer {
	v.userAgent = userAgent
	return *v
}

func (v *Viewer) UserAgent() string {
	return v.userAgent
}