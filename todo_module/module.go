package todo_module

import "echo-modarch/app"

type ToDoModule struct{}

func (tm *ToDoModule) RegisterRoutes() {
	app.AddApplicationRoute("/todos/:id", "GET", getExampleToDoPage, "todoPageExample")
}

func (tm *ToDoModule) RegisterTemplates() {
	app.TempRender.AddTemplate("todo_module/views/*.html")
}
