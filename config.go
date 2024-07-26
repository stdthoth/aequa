package aequa

type config struct {
	Port        string
	View        string
	sessionType string
	cookie      cookieConfig
}

type cookieConfig struct {
	name     string
	persists string
	lifetime string
	secure   string
	domain   string
}
