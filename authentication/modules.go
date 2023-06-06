package authentication

import "echo-modarch/app"

type AuthenticationModule struct{}

func (am *AuthenticationModule) RegisterRoutes() {
	app.AddApplicationRoute("/users", "POST", saveUser, "saveUser")
}

func (am *AuthenticationModule) RegisterTemplates() {
	// Nothing to do here
}
