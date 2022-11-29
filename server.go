package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	Id        string
	PageTitle string
	Todos     []Todo
}

func getUser(c echo.Context) error {
	id := c.Param("id") // User ID from path `users/:id`
	data := TodoPageData{
		Id:        id,
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	err := c.Render(http.StatusOK, "index.html", data)

	if err != nil {
		log.Println(err)
	}

	return err
}

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", saveUser).Name = "saveUser"
	e.GET("/users/:id", getUser).Name = "getUser"

	e.Static("/static", "static")

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.Logger.Fatal(e.Start(":1323"))
}
