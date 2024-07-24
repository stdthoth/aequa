package tests

import (
	"os"
	"testing"

	"github.com/CloudyKit/jet/v6"
	"github.com/stdthoth/aequa/views"
)

var view = jet.NewSet(
	jet.NewOSFileSystemLoader("./testdata/views"),
	jet.InDevelopmentMode(),
)

var testViewer = views.View{
	Viewer:   "",
	RootPath: "",
	JetViews: view,
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
