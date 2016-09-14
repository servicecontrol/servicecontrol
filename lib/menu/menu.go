package menu

import ()

var (
	menu Menu
)

type Menu struct {
	MenuItems []MenuItem `json:"MenuItems"`
}

type MenuItem struct {
	InternalName    string `json:"internalName"`
	PublicName      string `json:"publicName"`
	Route           string `json:"route"`
	Description     string `json:"description"`
	IsVisibleInMain bool   `json:"isVisibleInMain"`
}

func SetConfig(m Menu) {
	menu = m
}

func Config() Menu {
	return menu
}
