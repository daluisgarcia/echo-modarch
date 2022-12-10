package landing_page

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func indexView(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! Try to requestthe route: GET /todos/1")
}
