package authentication

import "github.com/daluisgarcia/echo-framework-modular-arquitecture/app"

type AuthenticationModule struct{}

func (am *AuthenticationModule) RegisterRoutes() {
	app.AddApplicationRoute("/users", "POST", saveUser, "saveUser")
}

func (am *AuthenticationModule) RegisterTemplates() {
	// Nothing to do here
}
