package boot

import (
"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"runtime"

"servicecontrol.io/servicecontrol/lib/router"
"servicecontrol.io/servicecontrol/lib/view"
"servicecontrol.io/servicecontrol/lib/session"
"servicecontrol.io/servicecontrol/lib/server"


"servicecontrol.io/servicecontrol/controller"
)

// Info contains the application settings.
type Info struct {
//	Asset      asset.Info    `json:"Asset"`
//	Email      email.Info    `json:"Email"`
//	Form       form.Info     `json:"Form"`
//	Generation generate.Info `json:"Generation"`
//	MySQL      mysql.Info    `json:"MySQL"`
	Server     server.Info   `json:"Server"`
	Session    session.Info  `json:"Session"`
	Template   view.Template `json:"Template"`
	View       view.Info     `json:"View"`
	Path       string
}

func (c *Info) ParseJSON(b []byte) error {
	return json.Unmarshal(, &c)
}


func init() {
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}


func LoadConfig(configFile string) *Info {
	config := &Info{}
	jsonconfig.LoadOrFatal(configFile, config)
	config.Path = configFile
	return config
}

func RegisterServices(config *Info) {
	session.SetConfig(config.Session)
	controller.LoadRoutes()
	view.SetConfig(config.View)
	view.SetTemplates(config.Template.Root, config.Template.Children)
}



// SetUpMiddleware contains the middleware that applies to every request.
func SetUpMiddleware(h http.Handler) http.Handler {
	return router.ChainHandler( // Chain middleware, top middlware runs first
		h,                    // Handler to wrap
	)
}
