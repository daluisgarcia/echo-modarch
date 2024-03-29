package authentication

import (
	"echo-modarch/app"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
	userData := new(RegisterUserRequest)

	if err := c.Bind(userData); err != nil {
		log.Printf("ERROR BINDING DATA REGISTERING USER: %v\n", err)
		// TODO Return beauty error
		return err
	}

	service := NewUserService()

	userRegistered, err := service.RegisterUser(c.Request().Context(), userData)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, userRegistered)
}

func loginForm(c echo.Context) error {
	data := map[string]interface{}{
		"id": "ID",
	}
	return c.Render(http.StatusOK, "login_form.html", data)
}

func loginUser(c echo.Context) error {
	userData := new(LoginUserRequest)

	if err := c.Bind(userData); err != nil {
		log.Printf("ERROR BINDING DATA LOGGING USER: %v\n", err)
		// TODO Return beauty error
		return err
	}

	service := NewUserService()
	userFromDb, err := service.LoginUser(c.Request().Context(), userData)

	if err != nil {
		return err
	}

	userCookie := &app.UserCookieData{
		Id:              userFromDb.Id,
		Name:            userFromDb.Name,
		Email:           userFromDb.Email,
		IsAuthenticated: true,
	}

	if err = app.SetSessionCookie(c, userCookie); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, userFromDb)

}

func logoutUser(c echo.Context) error {

	err := app.RemoveSessionCookie(c)

	if err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, c.Echo().Reverse("loginUserGet"))
}
