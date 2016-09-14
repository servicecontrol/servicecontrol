package services

import (
	"fmt"
	"net/http"
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

	view.ExtractPageInfo(v.Vars, menu.Config())
	fmt.Println(v)
	if session.Values["id"] != nil {
		v.Vars["first_name"] = session.Values["first_name"]
	}
	v.Render(w, r)
}
