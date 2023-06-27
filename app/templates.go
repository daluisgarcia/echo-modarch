package app

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates map[string]*template.Template
}

// Renders a template
func (tr *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if temp, ok := tr.templates[name]; ok {
		return temp.ExecuteTemplate(w, "base", data) // Renders the base template with the blocks defined in the template got by name
	}

	err := fmt.Errorf("ERROR RENDERING TEMPLATE: Template %v not found", name)
	fmt.Println(err)

	return err
}

// Base templates
var baseTemplates = []string{
	"views/base.html",
	"views/navbar.html",
}

// Add a template to the renderer
func (tr *TemplateRenderer) AddTemplate(path string) {
	pathSplitted := strings.Split(path, "/")
	templateName := pathSplitted[len(pathSplitted)-1]

	temp := template.New("")

	// Adding functions to templates
	temp.Funcs(
		template.FuncMap{
			"reverse": echoApp.Reverse,
		},
	)

	templatesBatch := append(baseTemplates, path)
	temp, err := temp.ParseFiles(templatesBatch...) // Parsing the base templates and the template got by path

	if err != nil {
		panic(fmt.Errorf("ERROR ADDING TEMPLATE PATH \"%v\": %v\n", path, err))
	}

	tr.templates[templateName] = template.Must(temp, err)

}

// Global template renderer
var TempRender = &TemplateRenderer{
	templates: map[string]*template.Template{},
}
