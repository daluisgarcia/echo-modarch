// IMPORTANT: THIS FILE SHOULD NOT BE EDITED

package app

import (
	"echo-modarch/database"
	"echo-modarch/utils"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var echoApp *echo.Echo // App instance

// Initilize the server, setting the renderer and static files folder
func InitServer() error {
	// Loading environment variables into the config struct
	if err := utils.SetConfig(); err != nil {
		return err
	}

	echoApp = echo.New()
	echoApp.Use(session.Middleware(sessions.NewCookieStore([]byte(utils.GetConfig().SecretKey))))

	echoApp.Renderer = TempRender
	echoApp.Static("/static", "static")

	databaseConnection, err := database.NewPostgresDatabase()

	if err != nil {
		return err
	}

	database.SetDatabaseConnection(databaseConnection)

	return nil
}

// Run the server
func RunServer() {
	echoApp.Logger.Fatal(echoApp.Start(":1323"))
}
