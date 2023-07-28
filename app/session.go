// IMPORTANT: THIS FILE CAN BE EDITED TO FIT YOUR NEEDS

package app

import (
	"fmt"
	"log"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type UserCookieData struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

// SetSessionCookie sets a session cookie with the user data
func SetSessionCookie(c echo.Context, userCookieData *UserCookieData) error {
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

// GetUserFromCookie gets the user data from the session cookie
func GetUserFromCookie(c echo.Context) (*UserCookieData, error) {
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

// RemoveSessionCookie removes the session cookie
func RemoveSessionCookie(c echo.Context) error {
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

// SetUserInContext sets the user data in the context
func SetUserInContext(c echo.Context, userCookie *UserCookieData) echo.Context {
	c.Set("user", userCookie)
	return c
}

// GetUserFromContext gets the user data from the context
func GetUserFromContext(c echo.Context) (*UserCookieData, error) {
	userCookie, ok := c.Get("user").(*UserCookieData)

	if !ok {
		log.Printf("ERROR GETTING USER FROM CONTEXT\n")
		return nil, fmt.Errorf("ERROR GETTING USER FROM CONTEXT")
	}

	return userCookie, nil
}
