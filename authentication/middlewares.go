package authentication

import (
	"echo-modarch/app"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserIsNotLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userCookie, err := app.GetUserFromCookie(c)

		if err != nil {
			return err
		}

		if userCookie.Id != "" {
			// User is logged in
			return c.Redirect(http.StatusFound, "/")
		}

		return next(c)
	}
}

func UserIsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userCookie, err := app.GetUserFromCookie(c)

		if err != nil {
			return err
		}

		if userCookie.Id != "" {
			// User is logged in

			// TODO - Check if user exists in database

			c = app.SetUserInContext(c, userCookie)
			return next(c)
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("loginUserGet"))
	}
}
