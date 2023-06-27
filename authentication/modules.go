package authentication

import (
	"echo-modarch/app"
	"encoding/gob"
)

type AuthenticationModule struct{}

func (am *AuthenticationModule) RegisterRoutes() {
	gob.Register(&UserCookieData{})
	app.AddApplicationRoute("/register", "GET", registerForm, "registerUserForm")
	app.AddApplicationRoute("/register", "POST", registerUser, "saveUser")
	app.AddApplicationRoute("/login", "GET", loginForm, "loginUserForm")
	app.AddApplicationRoute("/login", "POST", loginUser, "loginUser")
	app.AddApplicationRoute("/logout", "GET", logoutUser, "logoutUser")
}

func (am *AuthenticationModule) RegisterTemplates() {
	app.TempRender.AddTemplate("authentication/views/login_form.html")
	app.TempRender.AddTemplate("authentication/views/register_form.html")
}
