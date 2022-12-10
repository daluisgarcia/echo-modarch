package authentication

import "github.com/daluisgarcia/golang-echo-test/app"

type AuthenticationModule struct{}

func (am *AuthenticationModule) RegisterRoutes() {
	app.AddApplicationRoute("/users", "POST", saveUser, "saveUser")
}

func (am *AuthenticationModule) RegisterTemplates() {
	// Nothing to do here
}
