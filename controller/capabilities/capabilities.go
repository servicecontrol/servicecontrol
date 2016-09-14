package capabilities

import (
	"net/http"

	"servicecontrol.io/servicecontrol/lib/menu"
	"servicecontrol.io/servicecontrol/lib/router"
	"servicecontrol.io/servicecontrol/lib/session"
	"servicecontrol.io/servicecontrol/lib/view"
)

const (
	uri          string = "/capabilities"
	viewTemplate string = "capabilities/index"
	intName      string = "capabilities"
)

// Load loads all routers for capabilities
func Load() {
	router.Get(uri, Index)
}

// Index responds to all GET requests
func Index(w http.ResponseWriter, r *http.Request) {
	session := session.Instance(r)

	v := view.New(viewTemplate)
	v.Vars["int_name"] = intName

	if session.Values["id"] != nil {
		v.Vars["first_name"] = session.Values["first_name"]
	}
	v.Render(w, r)
}
