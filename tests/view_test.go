package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemplate(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/url", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	testViewer.Viewer = "go"
	testViewer.RootPath = "./tests/testdata"

	err = testViewer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("error rendering page", err)
	}

	err = testViewer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("error rendering a non-existing go template", err)
	}

	testViewer.Viewer = "jet"
	testViewer.RootPath = "./tests/testdata"
	err = testViewer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("error rendering page", err)
	}

	err = testViewer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("error rendering a non-existing jet template", err)
	}

	testViewer.Viewer = ""
	err = testViewer.Page(w, r, "home", nil, nil)
	if err == nil {
		t.Error("no error while returning a non exisitent template", err)
	}
}

func TestGoTemplate(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testViewer.Viewer = "go"
	testViewer.RootPath = "./tests/testdata"

	err = testViewer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("error rendering page", err)
	}
}

func TestJetTemplate(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testViewer.Viewer = "jet"
	testViewer.RootPath = "./tests/testdata"
	err = testViewer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error(err)
	}

}
