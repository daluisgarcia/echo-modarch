package main

import (
	"log"

	"github.com/daluisgarcia/golang-echo-test/app"
)

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
