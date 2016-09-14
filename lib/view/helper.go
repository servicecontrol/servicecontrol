package view

import (
	"servicecontrol.io/servicecontrol/lib/menu"
)

func ExtractPageInfo(viewConfig map[string]interface{}, m menu.Menu) {
	for _, v := range m.MenuItems {
		if v.InternalName == viewConfig["int_name"] {
			viewConfig["page_title"] = v.PublicName
			viewConfig["page_route"] = v.Route
			viewConfig["page_description"] = v.Description
			viewConfig["isVisibleInMain"] = v.IsVisibleInMain
		}
	}
}
