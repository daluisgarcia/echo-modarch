package authentication

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserIsNotLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userCookie, err := getSessionCookie(c)

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
		userCookie, err := getSessionCookie(c)

		if err != nil {
			return err
		}

		if userCookie.Id != "" {
			// User is logged in
			c.Set("user", userCookie)
			return next(c)
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("loginUserGet"))
	}
}
