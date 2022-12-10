package app

import "github.com/labstack/echo/v4"

// Allows to add a route to the echo app
func AddApplicationRoute(endpoint, httpMethod string, handler func(c echo.Context) error, routeName string) {
	switch httpMethod {
	case "GET":
		echoApp.GET(endpoint, handler).Name = routeName
	case "POST":
		echoApp.POST(endpoint, handler).Name = routeName
	case "PUT":
		echoApp.PUT(endpoint, handler).Name = routeName
	case "DELETE":
		echoApp.DELETE(endpoint, handler).Name = routeName
	case "PATCH":
		echoApp.PATCH(endpoint, handler).Name = routeName
	default:
		echoApp.GET(endpoint, handler).Name = routeName
	}
}
