package todo_module

import (
	"echo-modarch/app"
	"log"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ToDoService struct {
	repo *ToDoRepository
}

func NewToDoService() *ToDoService {
	return &ToDoService{
		repo: NewToDoRepository(),
	}
}

func (this *ToDoService) GetToDosList(c echo.Context) ([]*Task, error) {
	userLogged, err := app.GetUserFromContext(c)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return this.repo.GetAllTasks(c.Request().Context(), userLogged.Id)
}

func (this *ToDoService) GetTaskById(c echo.Context, id string) (*Task, error) {
	return this.repo.GetTaskById(c.Request().Context(), id)
}

func (this *ToDoService) CreateNewTask(c echo.Context, task *CreateTaskData) error {
	userLogged, err := app.GetUserFromContext(c)

	if err != nil {
		log.Println(err)
		return err
	}

	taskToInsert := Task{
		Id:          uuid.New().String(),
		Title:       task.Title,
		Description: task.Description,
	}

	return this.repo.InsertTask(c.Request().Context(), taskToInsert, userLogged.Id)
}
