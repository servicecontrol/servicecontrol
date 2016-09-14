package usage

import (
	"net/http"

	"servicecontrol.io/servicecontrol/lib/menu"
	"servicecontrol.io/servicecontrol/lib/router"
	"servicecontrol.io/servicecontrol/lib/session"
	"servicecontrol.io/servicecontrol/lib/view"
)

const (
	uri          string = "/usage"
	viewTemplate string = "usage/index"
	intName      string = "usage"
)

// Load loads all routes for usage
func Load() {
	router.Get(uri, Index)
}

// Index handles all GET requests
func Index(w http.ResponseWriter, r *http.Request) {
	session := session.Instance(r)

	v := view.New(viewTemplate)
	v.Vars["int_name"] = intName

	view.ExtractPageInfo(v.Vars, menu.Config())

	if session.Values["id"] != nil {
		v.Vars["first_name"] = session.Values["first_name"]
	}
	v.Render(w, r)
}
