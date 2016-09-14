package pageinfo

import (
	"net/http"

	"servicecontrol.io/servicecontrol/lib/menu"
	"servicecontrol.io/servicecontrol/lib/view"
)

// ExtractPageInfo extracts page info from app config and adds it to input map
// for page rendering
func ExtractPageInfo(w http.ResponseWriter, r *http.Request, v *view.Info) {
	for _, val := range menu.Config().Items {
		if val.InternalName == v.Vars["int_name"] {
			v.Vars["page_title"] = val.PublicName
			v.Vars["page_route"] = val.Route
			v.Vars["page_description"] = val.Description
			v.Vars["is_visible_in_main_menu"] = val.IsVisibleInMain
			v.Vars["menu_items"] = menu.Config().Items
		}
	}
}
