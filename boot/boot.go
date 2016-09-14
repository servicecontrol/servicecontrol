package boot

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"

	"servicecontrol.io/servicecontrol/lib/asset"
	"servicecontrol.io/servicecontrol/lib/jsonconfig"
	"servicecontrol.io/servicecontrol/lib/menu"
	"servicecontrol.io/servicecontrol/lib/router"
	"servicecontrol.io/servicecontrol/lib/server"
	"servicecontrol.io/servicecontrol/lib/session"
	"servicecontrol.io/servicecontrol/lib/view"

	"servicecontrol.io/servicecontrol/controller"
)

// AppConfig contains the application settings.
type AppConfig struct {
	Asset asset.Info `json:"Asset"`
	//	Email      email.Info    `json:"Email"`
	//	Form       form.Info     `json:"Form"`
	//	Generation generate.Info `json:"Generation"`
	//	MySQL      mysql.Info    `json:"MySQL"`
	Server   server.Info   `json:"Server"`
	Session  session.Info  `json:"Session"`
	Template view.Template `json:"Template"`
	View     view.Info     `json:"View"`
	Menu     menu.Menu     `json:"Menu"`
	Path     string
}

// ParseJSON parses JSON into a Config struct
func (c *AppConfig) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

func init() {
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// LoadConfig loads all app config from a file
func LoadConfig(configFile string) *AppConfig {
	config := &AppConfig{}
	jsonconfig.LoadOrFatal(configFile, config)
	config.Path = configFile
	return config
}

// RegisterServices passes app config to the according service(lib)
func RegisterServices(config *AppConfig) {
	session.SetConfig(config.Session)
	menu.SetConfig(config.Menu)
	controller.LoadRoutes()
	asset.SetConfig(config.Asset)
	view.SetConfig(config.View)
	view.SetTemplates(config.Template.Root, config.Template.Children)
	// Set up the functions for the views
	view.SetFuncMaps(
		asset.Map(config.View.BaseURI),
	)
}

// SetUpMiddleware contains the middleware that applies to every request.
func SetUpMiddleware(h http.Handler) http.Handler {
	return router.ChainHandler( // Chain middleware, top middlware runs first
		h, // Handler to wrap
	)
}
