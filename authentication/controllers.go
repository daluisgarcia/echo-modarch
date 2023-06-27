package authentication

import (
	"echo-modarch/utils"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func getSessionCookie(c echo.Context) (*UserCookieData, error) {
	loginCookie, err := session.Get("login", c)

	if err != nil {
		log.Printf("ERROR GETTING COOKIE: %v\n", err)
		return nil, err
	}

	cookieUser := loginCookie.Values["user"]
	if userLogged, ok := cookieUser.(*UserCookieData); ok {
		return userLogged, nil
	}

	return &UserCookieData{}, nil
}

func setSessionCookie(c echo.Context, userCookieData *UserCookieData) error {
	loginCookie, err := session.Get("login", c)

	if err != nil {
		log.Printf("ERROR GETTING COOKIE: %v\n", err)
		return err
	}

	loginCookie.Options = &sessions.Options{
		Path:     "/login",
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

func removeSessionCookie(c echo.Context) error {
	loginCookie, err := session.Get("login", c)

	if err != nil {
		log.Printf("ERROR GETTING COOKIE: %v\n", err)
		return err
	}

	loginCookie.Options = &sessions.Options{
		Path:     "/login",
		MaxAge:   -1,
		HttpOnly: true,
	}
	loginCookie.Values["user"] = nil

	return loginCookie.Save(c.Request(), c.Response())
}

func registerForm(c echo.Context) error {
	data := map[string]interface{}{
		"id": "ID",
	}
	err := c.Render(http.StatusOK, "register_form.html", data)

	if err != nil {
		log.Println(err)
	}

	return err
}

func registerUser(c echo.Context) error {
	userData := new(User)

	if err := c.Bind(userData); err != nil {
		// TODO Return beauty error
		return err
	}
	userData.Id = uuid.New().String()

	var err error
	userData.Password, err = utils.HashString(userData.Password)

	if err != nil {
		// TODO Return beauty handlded error
		return err
	}

	return c.JSON(http.StatusCreated, userData)
}

func loginForm(c echo.Context) error {
	userLogged, err := getSessionCookie(c)

	if userLogged.Id == "" && err != nil {
		return err
	}

	if userLogged.Id != "" {
		// User is logged in
		return c.JSON(http.StatusOK, userLogged)
	}

	data := map[string]interface{}{
		"id":       "ID",
		"loginUrl": c.Echo().Reverse("loginUser"),
	}
	return c.Render(http.StatusOK, "login_form.html", data)
}

func loginUser(c echo.Context) error {
	userLogged, err := getSessionCookie(c)

	if userLogged.Id == "" && err != nil {
		return err
	}

	if userLogged.Id != "" {
		// User is logged in
		return c.JSON(http.StatusOK, userLogged)
	}

	userData := new(LoginUserRequest)

	if err := c.Bind(userData); err != nil {
		log.Printf("ERROR BINDING DATA: %v\n", err)
		return err
	}

	// TODO Get user from database

	userLogged.Email = userData.Email
	userLogged.Id = uuid.New().String()

	if err = setSessionCookie(c, userLogged); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, userLogged)

}

func logoutUser(c echo.Context) error {

	err := removeSessionCookie(c)

	if err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, c.Echo().Reverse("loginUser"))
}
