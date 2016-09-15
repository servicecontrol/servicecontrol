package management

import (
	"net/http"

	"servicecontrol.io/servicecontrol/lib/router"
	"servicecontrol.io/servicecontrol/lib/session"
	"servicecontrol.io/servicecontrol/lib/view"
)

const (
	uri          string = "/management"
	viewTemplate string = "management/index"
)

// Load configures all routers
func Load() {
	router.Get(uri, Index)
}

// Index responds to all GET requests
func Index(w http.ResponseWriter, r *http.Request) {
	session := session.Instance(r)

	v := view.New(viewTemplate)

	if session.Values["id"] != nil {
		v.Vars["first_name"] = session.Values["first_name"]
	}
	v.Render(w, r)
}
