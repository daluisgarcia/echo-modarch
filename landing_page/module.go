package landing_page

import "echo-modarch/app"

type LandingModule struct{}

func (tm *LandingModule) RegisterRoutes() {
	app.AddApplicationRoute("/", "GET", indexView, "indexView")
}

func (tm *LandingModule) RegisterTemplates() {
	// not implemented
}
