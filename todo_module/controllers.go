package todo_module

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getExampleToDoPage(c echo.Context) error {
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
	err := c.Render(http.StatusOK, "todo_index.html", data) // To reference the template file inside the folder, it must have an unique name

	if err != nil {
		log.Println(err)
	}

	return err
}
