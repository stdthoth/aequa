package views

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet/v6"
)

type View struct {
	Viewer     string
	JetViews   *jet.Set
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
}

func (v *View) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(v.Viewer) {
	case "go":
		return v.GoTemplates(w, r, view, data)
	case "jet":
		return v.JetTemlates(w, r, view, variables, data)
	default:

	}

	return errors.New("no rendering engine specified")
}

// GoTemplates renders standard go.html templates
func (v *View) GoTemplates(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf(".%s/views/%s.page.tmpl", v.RootPath, view))
	if err != nil {
		return err
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	err = tmpl.Execute(w, &td)
	if err != nil {
		return err
	}

	return nil
}

// JetTemplate renders a template using the Jet package
func (v *View) JetTemlates(w http.ResponseWriter, r *http.Request, tmplName string, variables, data interface{}) error {
	var JetVar jet.VarMap

	if variables == nil {
		JetVar = make(jet.VarMap)
	} else {
		JetVar = variables.(jet.VarMap)
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	tmpl, err := v.JetViews.GetTemplate(fmt.Sprintf("%s.jet", tmplName))
	if err != nil {
		log.Println(err)
		return err
	}

	if err = tmpl.Execute(w, JetVar, td); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
