package dashboard

import (
	"net/http"
	"servicecontrol.io/servicecontrol/lib/router"
	"servicecontrol.io/servicecontrol/lib/session"
	"servicecontrol.io/servicecontrol/lib/view"
)

func Load() {
	router.Get("/", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	session := session.Instance(r)

	v := view.New("dashboard/index")
	v.Vars["page_title"] = "Dashboard"
	
	if session.Values["id"] != nil {
		v.Vars["first_name"] = session.Values["first_name"]
	}
	v.Render(w, r)
}
