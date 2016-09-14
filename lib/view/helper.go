package view

import (
	"servicecontrol.io/servicecontrol/lib/menu"
)

// ExtractPageInfo extracts page info from app config and adds it to iput map for page rendering
func ExtractPageInfo(viewConfig map[string]interface{}, m menu.Menu) {
	for _, v := range m.MenuItems {
		if v.InternalName == viewConfig["int_name"] {
			viewConfig["page_title"] = v.PublicName
			viewConfig["page_route"] = v.Route
			viewConfig["page_description"] = v.Description
			viewConfig["is_visible_in_main_menu"] = v.IsVisibleInMain
			viewConfig["menu_items"] = m.MenuItems
		}
	}
}
