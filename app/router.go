// IMPORTANT: THIS FILE SHOULD NOT BE EDITED

package app

import "github.com/labstack/echo/v4"

// Allows to add a route to the echo app providing the endpoint, the http method,
// the handler, the route name and the middlewares applied to that route
func AddApplicationRoute(
	endpoint,
	httpMethod string,
	handler func(c echo.Context) error,
	routeName string,
	middlewares ...echo.MiddlewareFunc,
) {
	switch httpMethod {
	case "GET":
		echoApp.GET(endpoint, handler, middlewares...).Name = routeName
	case "POST":
		echoApp.POST(endpoint, handler, middlewares...).Name = routeName
	case "PUT":
		echoApp.PUT(endpoint, handler, middlewares...).Name = routeName
	case "DELETE":
		echoApp.DELETE(endpoint, handler, middlewares...).Name = routeName
	case "PATCH":
		echoApp.PATCH(endpoint, handler, middlewares...).Name = routeName
	default:
		echoApp.GET(endpoint, handler, middlewares...).Name = routeName
	}
}
