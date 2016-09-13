package account

import (
	"net/http"
	"servicecontrol.io/servicecontrol/lib/router"
	"servicecontrol.io/servicecontrol/lib/session"
	"servicecontrol.io/servicecontrol/lib/view"
	"servicecontrol.io/servicecontrol/lib/menu"
)

func Load() {
	router.Get("/account", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	session := session.Instance(r)

	v := view.New("account/index")
	v.Vars["int_name"] = "Account"
	v.Vars["menu_items"] = menu.Config().MenuItems

        view.ExtractPageInfo(v.Vars, menu.Config())

	if session.Values["id"] != nil {
		v.Vars["first_name"] = session.Values["first_name"]
	}
	v.Render(w, r)
}
