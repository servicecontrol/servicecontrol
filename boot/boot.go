package boot

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/context"

	"servicecontrol.io/servicecontrol/lib/asset"
	"servicecontrol.io/servicecontrol/lib/jsonconfig"
	"servicecontrol.io/servicecontrol/lib/menu"
	"servicecontrol.io/servicecontrol/lib/router"
	"servicecontrol.io/servicecontrol/lib/server"
	"servicecontrol.io/servicecontrol/lib/session"
	"servicecontrol.io/servicecontrol/lib/view"
	"servicecontrol.io/servicecontrol/viewmodify"

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

// init sets runtime settings.
func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// LoadConfig loads all app config from a file
func LoadConfig(configFile string) *AppConfig {
	config := &AppConfig{}
	jsonconfig.LoadOrFatal(configFile, config)

	//Store path to config file
	config.Path = configFile

	return config
}

// RegisterServices passes app config to the according service(lib)
func RegisterServices(config *AppConfig) {
	// Set up the session cookie store
	session.SetConfig(config.Session)

	// Set up CSRF protection
	// xsrf.SetConfig(xsrf.Info{
	// 	AuthKey: config.Session.CSRFKey,
	// 	Secure:  config.Session.Options.Secure,
	// })

	// Connect to database
	// mysql.SetConfig(config.MySQL)
	// mysql.Connect(true)

	// Configure form handling
	//form.SetConfig(config.Form)

	// Load the controller routes
	controller.LoadRoutes()

	// Set up the assets
	asset.SetConfig(config.Asset)

	// Set up the views
	view.SetConfig(config.View)
	view.SetTemplates(config.Template.Root, config.Template.Children)

	// Set up the functions for the views
	view.SetFuncMaps(
		asset.Map(config.View.BaseURI),
	//	link.Map(config.View.BaseURI),
	// noescape.Map(),
	// prettytime.Map(),
	// form.Map(),
	)

	// Set up the variables and modifiers for the views
	view.SetModifiers(
		// authlevel.Modify,
		// uri.Modify,
		// xsrf.Token,
		// flash.Modify,
		viewmodify.ExtractPageInfo,
	)
}

// SetUpMiddleware contains the middleware that applies to every request.
func SetUpMiddleware(h http.Handler) http.Handler {
	return router.ChainHandler( // Chain middleware, top middlware runs first
		h, // Handler to wrap
		//	setUpCSRF, // Prevent CSRF
		// rest.Handler,         // Support changing HTTP method sent via query string
		// logrequest.Handler,   // Log every request
		context.ClearHandler, // Prevent memory leak with gorilla.sessions
	)
}

// setUpCSRF sets up the CSRF protection
// func setUpCSRF(h http.Handler) http.Handler {
// 	// Decode the string
// 	key, err := base64.StdEncoding.DecodeString(xsrf.Config().AuthKey)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// Configure the middleware
// 	cs := csrf.Protect([]byte(key),
// 		csrf.ErrorHandler(http.HandlerFunc(status.InvalidToken)),
// 		csrf.FieldName("_token"),
// 		csrf.Secure(xsrf.Config().Secure),
// 	)(h)
// 	return cs
// }
