package authentication

import (
	"echo-modarch/app"
	"encoding/gob"
)

type AuthenticationModule struct{}

func (am *AuthenticationModule) RegisterRoutes() {
	gob.Register(&app.UserCookieData{})
	app.AddApplicationRoute("/register", "GET", registerForm, "registerUserForm", UserIsNotLoggedIn)
	app.AddApplicationRoute("/register", "POST", registerUser, "saveUser", UserIsNotLoggedIn)
	app.AddApplicationRoute("/login", "GET", loginForm, "loginUserGet", UserIsNotLoggedIn)
	app.AddApplicationRoute("/login", "POST", loginUser, "loginUserPost", UserIsNotLoggedIn)
	app.AddApplicationRoute("/logout", "GET", logoutUser, "logoutUser", UserIsLoggedIn)
}

func (am *AuthenticationModule) RegisterTemplates() {
	app.TempRender.AddTemplate("authentication/views/login_form.html")
	app.TempRender.AddTemplate("authentication/views/register_form.html")
}
