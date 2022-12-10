package todo_module

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	Id        string
	PageTitle string
	Todos     []Todo
}