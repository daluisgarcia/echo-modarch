# Golang Echo Framework Modular Base Architecture

This project defines a way to use the [Golang Echo Framework](https://echo.labstack.com/) as a SSR (Server Side Rendering) framework based in [Laravel](https://laravel.com/) abd [Django](https://www.djangoproject.com/), imitating its template usage and inheritance.

The idea of this base architecture is to provide a folder structure that allows developers to build faster a fullstack project with Echo. Defining a way to create modules that can be plugged or unplugged from the main server only by adding them to a slice.

This project is aimed at developers with a basic knowledge of Golang web servers and Echo Framework. If you are new to Golang or Echo, please read the [Golang documentation](https://golang.org/doc/) and the [Echo documentation](https://echo.labstack.com/docs) first.

## Project structure

This project  is structured in the following way:

- `app/` This folder contains the server and database driver instantiation. Also contains the code to lift the modules and subcribe the routes and controlles to the main server. **THE FILES INSIDE THIS FOLDER SHOULD NOT BE MODIFIED**.

- `authentication/` This folder contains the code to manage the creation and authentication of the users. The session is handled by a session cookie and the user data is added to the context by the UserIsLoggedIn middleware. Inside the `views/` folder you can see how template ingeritance and template custom functions can be used. **THE CODE INSIDE THIS FOLDER IS A WORKING EXAMPLE OF HOW AUTHENTICATION CAN BE IMPLEMENTED IN ECHO**, the code may be modified if it doesn't fit your project requirements.

- `database/` This folder contains the code to connect to the database and a struct to be used in compositions with other repositories structs in other modules. Also contains a `up.sql` file that allows to initialize the database. **THIS CODE CAN BE EDITED TO FIT YOUR PROJECT REQUIREMENTS.** An ORM can be used instead of the database/sql package.

- `landing_page/` This folder contains the minimum code a module needs to work correctly. This module only renders a string and the `modules.go` file contains the code to subscribe the routes and controllers to the main server. **THIS CODE CAN BE EDITED TO FIT YOUR PROJECT REQUIREMENTS.**

- `todo_modules/` This folder contains a module to manage a todo list. This module is a working example of how a module can be created and plugged to the main server. Inside the `views/` folder you can see how template ingeritance can be used. **THIS CODE IS JUST AND EXAMPLE, IT CAN BE EDITED OR DELETED**

- `utils/` This folder contains some utilities as the project configuration management logic and struct (used in database instantiation) and hasher methods (used in password hashing). **THIS CODE CAN BE EDITED TO FIT YOUR PROJECT REQUIREMENTS.**

- `views/` This folder contains the base html templates used in the project. THe templates here are used in template inheritance by the server renderer. **If you need to add/remove some base template to the project, please view the `app/templates.go` file first.**

- `main.go` This file contains the code to run the server. **This code should be edited only to add modules to the `modulesToRegister` slice**.

## How to create a new module

As it was said before, the `landing_page/` folder contains the minimum code necessary in a module. To create a module you need a struct that implements the `IAppModule` interface defined in the `app/modules.go` file. This interface defines the methods that a module should implement to be plugged to the main server. Theinterface is:

    type IAppModule interface {
        RegisterRoutes()
        RegisterTemplates()
    }

The `RegisterRoutes()` method should contain the code to subscribe the routes and controllers to the main server. The `RegisterTemplates()` method should contain the code to subscribe the templates to the main server. The `todo_module/modules.go` file is a good example of how to implement this methods.

When the module struct is created, you only need to add it to the `modulesToRegister` slice in the `main.go` file. After doing this your module will be plugged to the main server and will be accesible in the routes you defined in the `RegisterRoutes()` method.

**This architecture can be used to create an API too.** You only need to use methods as `c.JSON()` defined by Echo. In this case you don't need to implemnt the `RegisterTemplates()` method of your module struct, just define it as an empty method to fulfill the `IAppModule` interface.

## Install packages

To install the project packages run the command `go get` in the root directory of the project.

## Run the project

First you need to create a `.env` as a copy of the `.env.example` file and fulfill it with values for the variables. After doing this, you just need to run the command `go run main.go` in the root directory of the project to execute the project.

## I want to contribute to the project

If you want to contribute to the project, please feel free to fork the project and create a pull request with your changes. I will review it and add it to the project if it fits the project vision.

If you want to contact me, you can do it by my [LinkedIn](https://www.linkedin.com/in/daluisgarcia/)
