package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

//LayoutDir a
var ViewsDir string = "views"
var ViewsExt string = ".gohtml"
var LayoutDir string = ViewsDir + "/layouts"

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*" + ViewsExt)
	if err != nil {
		panic(err)
	}
	return files
}

//NewView sa
func NewView(layout string, files ...string) *View {
	addViewsDirPrefix(files)
	addViewExtSuffix(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

//View 	a
type View struct {
	Template *template.Template
	Layout   string
}

//Render w
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)
}

func addViewsDirPrefix(files []string) {
	for i, f := range files {
		files[i] = ViewsDir + "/" + f
	}
}

func addViewExtSuffix(files []string) {
	for i, f := range files {
		files[i] = f + ViewsExt
	}
}
