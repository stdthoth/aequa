package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Session struct {
	Name        string
	Lifetime    string
	Persists    string
	Domain      string
	Secure      string
	SessionType string
}

func (s *Session) NewSession() *scs.SessionManager {
	var secure, persists bool

	exp, err := strconv.Atoi(s.Lifetime)
	if err != nil {
		exp = 60
	}

	if strings.ToLower(s.Persists) == "true" {
		persists = true
	}

	if strings.ToLower(s.Secure) == "true" {
		secure = true
	}

	session := scs.New()
	session.Lifetime = time.Duration(exp) * time.Minute
	session.Cookie.Secure = secure
	session.Cookie.Persist = persists
	session.Cookie.Name = s.Name
	session.Cookie.Domain = s.Domain
	session.Cookie.SameSite = http.SameSiteLaxMode

	switch strings.ToLower(s.SessionType) {

	case "redis":

	case "memcache":

	case "postgresql", "postgres":

	case "mysql", "mariadb":

	case "sqlite3":

	default:

	}

	return session
}
