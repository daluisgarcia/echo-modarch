package app

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (tr *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return tr.templates.ExecuteTemplate(w, name, data)
}

func (tr *TemplateRenderer) AddTemplate(path string) {
	// Add the new templates to the existing ones
	tr.templates.ParseGlob(path)
}

// Global template renderer
var TempRender = &TemplateRenderer{
	templates: template.Must(template.ParseGlob("views/*.html")), // Defines the base templates folder
}
