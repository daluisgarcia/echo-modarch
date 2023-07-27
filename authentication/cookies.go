package authentication

import (
	"log"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func setSessionCookie(c echo.Context, userCookieData *UserCookieData) error {
	loginCookie, err := session.Get("login", c)

	if err != nil {
		log.Printf("ERROR GETTING COOKIE: %v\n", err)
		return err
	}

	loginCookie.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	}
	loginCookie.Values["user"] = userCookieData

	err = loginCookie.Save(c.Request(), c.Response())

	if err != nil {
		log.Printf("ERROR SAVING COOKIE: %v\n", err)
		return err
	}

	return nil
}

func getSessionCookie(c echo.Context) (*UserCookieData, error) {
	loginCookie, err := session.Get("login", c)

	if err != nil {
		log.Printf("ERROR GETTING COOKIE: %v\n", err)
		return nil, err
	}

	cookieUser := loginCookie.Values["user"]
	userLogged, ok := cookieUser.(*UserCookieData)

	if !ok {
		// A problem parsing ocurred, the user cookie must not be set
		return &UserCookieData{}, nil
	}

	return userLogged, nil
}

func removeSessionCookie(c echo.Context) error {
	loginCookie, err := session.Get("login", c)

	if err != nil {
		log.Printf("ERROR GETTING COOKIE: %v\n", err)
		return err
	}

	loginCookie.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	loginCookie.Values["user"] = nil

	return loginCookie.Save(c.Request(), c.Response())
}
