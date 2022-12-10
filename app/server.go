package app

import (
	"github.com/labstack/echo/v4"
)

var echoApp *echo.Echo // App instance

func InitServer() error {
	// Loading environment variables into the config struct
	if err := SetConfig(); err != nil {
		return err
	}

	echoApp = echo.New()

	echoApp.Renderer = TempRender
	echoApp.Static("/static", "static")

	return nil
}

func RunServer() {
	// Start server
	echoApp.Logger.Fatal(echoApp.Start(":1323"))
}
