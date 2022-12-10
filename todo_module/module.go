package todo_module

import "github.com/daluisgarcia/golang-echo-test/app"

type ToDoModule struct{}

func (tm *ToDoModule) RegisterRoutes() {
	app.AddApplicationRoute("/todos/:id", "GET", getExampleToDoPage, "todoPageExample")
}

func (tm *ToDoModule) RegisterTemplates() {
	app.TempRender.AddTemplate("todo_module/views/*.html")
}
