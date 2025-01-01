package aequa

import "net/http"

func (a Aequa) LoadAndSave(next http.Handler) http.Handler {
	return a.Session.LoadAndSave(next)
}