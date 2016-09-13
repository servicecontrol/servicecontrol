package services

import (
	"net/http"
	"fmt"
	"servicecontrol.io/servicecontrol/lib/menu"
	"servicecontrol.io/servicecontrol/lib/router"
	"servicecontrol.io/servicecontrol/lib/session"
	"servicecontrol.io/servicecontrol/lib/view"
)

const (
	uri string = "/services"
)

func Load() {
	router.Get(uri, Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	session := session.Instance(r)

	v := view.New("services/index")
	v.Vars["int_name"] = "services"
	v.Vars["menu_items"] = menu.Config().MenuItems
	
	extractPageInfo(v.Vars, menu.Config())	

	if session.Values["id"] != nil {
		v.Vars["first_name"] = session.Values["first_name"]
	}
	v.Render(w, r)
}

func extractPageInfo(viewConfig map[string]interface{}, m menu.Menu) {
	for k,v := range m.MenuItems {
		fmt.Println(v,k)
		if v.InternalName == viewConfig["int_name"] {
			viewConfig["page_title"] = v.PublicName
			viewConfig["route"] = v.Route
		}
	}
}
