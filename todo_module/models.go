package todo_module

type Task struct {
	Id          string `query:"id"`
	Title       string `query:"title"`
	Description string `query:"description"`
	Done        bool   `query:"done"`
}
