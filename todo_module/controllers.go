package todo_module

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func newTaskForm(c echo.Context) error {
	return c.Render(http.StatusOK, "create_task.html", make(map[string]interface{}))
}

func createNewTask(c echo.Context) error {
	taskData := new(CreateTaskData)

	if err := c.Bind(taskData); err != nil {
		log.Println(err)
		return err
	}

	service := NewToDoService()
	err := service.CreateNewTask(c, taskData)

	if err != nil {
		log.Println(err)
	}

	return c.Redirect(http.StatusFound, "/todos")
}

func getTaskList(c echo.Context) error {
	service := NewToDoService()
	taskList, err := service.GetToDosList(c)

	var errorString string
	if err != nil {
		errorString = err.Error()
		log.Println(err)
	}

	// Structured way to send context to the template
	data := TodoPageDTO{
		PageTitle: "My TODO list",
		Tasks:     taskList,
		Error:     errorString,
	}
	// To reference the template file inside the folder, it must have an unique name into the whole project
	err = c.Render(http.StatusOK, "task_list.html", data)

	if err != nil {
		log.Println(err)
	}

	return err
}

func taskDetails(c echo.Context) error {
	id := c.Param("id")

	service := NewToDoService()
	task, err := service.GetTaskById(c, id)

	var errorString string
	if err != nil {
		errorString = err.Error()
		log.Println(err)
	}

	data := map[string]interface{}{
		"task":  task,
		"error": errorString,
	}

	return c.Render(http.StatusOK, "task_detail.html", data)
}
