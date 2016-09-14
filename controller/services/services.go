package services

import (
	"net/http"

	"servicecontrol.io/servicecontrol/lib/router"
	"servicecontrol.io/servicecontrol/lib/session"
	"servicecontrol.io/servicecontrol/lib/view"
)

const (
	uri          string = "/services"
	viewTemplate string = "services/index"
)

// Load configures all routers for services
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
