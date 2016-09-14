package pageinfo

import (
	"net/http"

	"servicecontrol.io/servicecontrol/lib/menu"
	"servicecontrol.io/servicecontrol/lib/view"
)

// Modify extracts page info from app config and adds it to input map
// for page rendering
func Modify(w http.ResponseWriter, r *http.Request, v *view.Info) {
	for _, val := range menu.Config().Items {
		if val.Route == v.Vars["CurrentURI"] {
			v.Vars["page_title"] = val.PublicName
			v.Vars["page_route"] = val.Route
			v.Vars["page_description"] = val.Description
			v.Vars["is_visible_in_main_menu"] = val.IsVisibleInMain
			v.Vars["menu_items"] = menu.Config().Items
		}
	}
}
