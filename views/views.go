package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

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

//NewView will add the prefixes and suffixes and return a View
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

//View, This is a View
type View struct {
	Template *template.Template
	Layout   string
}

//Render Fucntion is required by gorillamux packages
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

//ServeHTTP function requried by gorillamux Package...capitalization is important :/
func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)
}

//Function to add directory prefixes so i dont have to manually from now on
func addViewsDirPrefix(files []string) {
	for i, f := range files {
		files[i] = ViewsDir + "/" + f
	}
}

//Function to add .gohtml extensions so I don't have to manually
func addViewExtSuffix(files []string) {
	for i, f := range files {
		files[i] = f + ViewsExt
	}
}
