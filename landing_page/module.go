package landing_page

import "github.com/daluisgarcia/golang-echo-test/app"

type LandingModule struct{}

func (tm *LandingModule) RegisterRoutes() {
	app.AddApplicationRoute("/", "GET", indexView, "indexView")
}

func (tm *LandingModule) RegisterTemplates() {
	// not implemented
}
