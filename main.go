package main

import (
	"log"

	"echo-modarch/app"
	"echo-modarch/authentication"
	"echo-modarch/landing_page"
	"echo-modarch/todo_module"
)

// Slice of the modules to be registered
// This is the place where you register your modules
var modulesToRegister = []app.IAppModule{
	&landing_page.LandingModule{},
	&todo_module.ToDoModule{},
	&authentication.AuthenticationModule{},
}

func main() {
	// Initializing the server
	err := app.InitServer()
	if err != nil {
		log.Fatalf("%v", err)
	}

	var modulesRegister app.ModuleRegister = app.ModuleRegister{}
	// Setting the modules to register
	modulesRegister.SetAppModules(modulesToRegister)
	// Lifting modules services, repositories and routes
	modulesRegister.LiftModules()
	// Running the server
	app.RunServer()
}
