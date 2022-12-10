package main

import (
	"github.com/daluisgarcia/golang-echo-test/app"
	"github.com/daluisgarcia/golang-echo-test/landing_page"
	"github.com/daluisgarcia/golang-echo-test/todo_module"
)

var modulesToRegister = []app.IAppModule{
	&landing_page.LandingModule{},
	&todo_module.ToDoModule{},
}
