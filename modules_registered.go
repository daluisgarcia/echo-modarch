package main

import (
	"github.com/daluisgarcia/echo-framework-modular-arquitecture/app"
	"github.com/daluisgarcia/echo-framework-modular-arquitecture/authentication"
	"github.com/daluisgarcia/echo-framework-modular-arquitecture/landing_page"
	"github.com/daluisgarcia/echo-framework-modular-arquitecture/todo_module"
)

var modulesToRegister = []app.IAppModule{
	&landing_page.LandingModule{},
	&todo_module.ToDoModule{},
	&authentication.AuthenticationModule{},
}
