package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string
	viewer        string
	template      string
	errorExpected bool
	errorMsg      string
}{
	{"go_tmpl", "go", "home", false, "error rendering go template"},
	{"go_tmpl_non-existent", "go", "no-file", true, "error rendering non existent template"},
	{"jet_tmpl", "jet", "home", false, "error rendering jet template"},
	{"jet_tmpl_non-existent", "jet", "no-file", true, "error rendering a non existent jet template"},
	{"jet_tmpl", "jet", "no-file", true, "error rendering a non existent jet template"},
}

func TestTemplate(t *testing.T) {

	for _, e := range pageData {
		r, err := http.NewRequest(http.MethodGet, "/url", nil)
		if err != nil {
			t.Error(err)
		}
		w := httptest.NewRecorder()
		testViewer.Viewer = e.viewer
		testViewer.RootPath = "./tests/testdata"

		err = testViewer.Page(w, r, "home", nil, nil)
		if e.errorExpected {
			if err != nil {
				t.Errorf("%s:%s", e.name, e.errorMsg)
			}
		} else {
			if err != nil {
				t.Errorf("%s:%s:%s", e.name, e.errorMsg, err.Error())
			}
		}
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
