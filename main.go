// IMPORTANT: THIS FILE SHOULD NOT BE EDITED

package main

import (
	"log"

	"echo-modarch/app"
	"echo-modarch/authentication"
	"echo-modarch/landing_page"
	"echo-modarch/todo_module"
)

// Slice that contains the modules to be registered.
// Modules to work with the server must be added here.
var modulesToRegister = []app.IAppModule{
	&authentication.AuthenticationModule{},
	&landing_page.LandingModule{},
	&todo_module.ToDoModule{},
}

// Project main function
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
