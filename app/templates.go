// IMPORTANT: THIS FILE SHOULD NOT BE EDITED

package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates map[string]*template.Template
}

// Renders a template
func (tr *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if temp, ok := tr.templates[name]; ok {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		var v interface{}
		if err := json.Unmarshal(dataBytes, &v); err != nil {
			log.Fatal(err)
		}

		if viewContext, isMap := v.(map[string]interface{}); isMap {
			var err error
			// If the context has a user, it will be passed to the template
			if viewContext["user"], err = GetUserFromCookie(c); err != nil {
				viewContext["user"] = &UserCookieData{}
			}

			// Renders the base template with the blocks defined in the template gotten by name
			return temp.ExecuteTemplate(w, "base", viewContext)
		}

		return fmt.Errorf("ERROR RENDERING TEMPLATE: Context cannot be parsed to map[string]interface{}")
	}

	err := fmt.Errorf("ERROR RENDERING TEMPLATE: Template %v not found", name)
	log.Printf("%v\n", err)

	return err
}

// Base templates used in templates rendering and extending
var baseTemplates = []string{
	"views/base.html",
	"views/navbar.html",
}

// Adds a template to the renderer
func (tr *TemplateRenderer) AddTemplate(path string) {
	pathSplitted := strings.Split(path, "/")
	templateName := pathSplitted[len(pathSplitted)-1]

	temp := template.New("")

	// Adding custom functions to templates
	temp.Funcs(
		template.FuncMap{
			"reverse": echoApp.Reverse,
		},
	)

	templatesBatch := append(baseTemplates, path)
	// Parsing the base templates and the template got by path
	temp, err := temp.ParseFiles(templatesBatch...)

	if err != nil {
		panic(fmt.Errorf("ERROR ADDING TEMPLATE PATH \"%v\": %v\n", path, err))
	}

	tr.templates[templateName] = template.Must(temp, err)

}

// Global template renderer
var TempRender = &TemplateRenderer{
	templates: map[string]*template.Template{},
}
