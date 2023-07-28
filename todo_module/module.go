package todo_module

import (
	"echo-modarch/app"
	"echo-modarch/authentication"
)

type ToDoModule struct{}

func (tm *ToDoModule) RegisterRoutes() {
	app.AddApplicationRoute("/todos", "GET", getTaskList, "tasksList", authentication.UserIsLoggedIn)
	app.AddApplicationRoute("/todos/new", "GET", newTaskForm, "newTaskForm", authentication.UserIsLoggedIn)
	app.AddApplicationRoute("/todos/new", "POST", createNewTask, "createNewTask", authentication.UserIsLoggedIn)
	app.AddApplicationRoute("/todos/:id", "GET", taskDetails, "taskDetails", authentication.UserIsLoggedIn)
}

func (tm *ToDoModule) RegisterTemplates() {
	app.TempRender.AddTemplate("todo_module/views/task_list.html")
	app.TempRender.AddTemplate("todo_module/views/create_task.html")
	app.TempRender.AddTemplate("todo_module/views/task_detail.html")
}
