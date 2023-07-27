// IMPORTANT: THIS FILE SHOULD NOT BE EDITED

package app

// Interface to define a module possible actions
type IAppModule interface {
	RegisterRoutes()
	RegisterTemplates()
}

type ModuleRegister struct {
	appModules []IAppModule
}

// Set the modules to register
func (m *ModuleRegister) SetAppModules(modules []IAppModule) {
	m.appModules = modules
}

// Lift all the modules registered services, repositories and routes
func (m *ModuleRegister) LiftModules() {
	for _, f := range m.appModules {
		f.RegisterTemplates()
		f.RegisterRoutes()
	}
}
