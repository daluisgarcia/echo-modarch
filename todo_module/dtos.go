package todo_module

type CreateTaskData struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}

type TodoPageDTO struct {
	PageTitle string
	Tasks     []*Task
	Error     string
}
