package menu

var (
	menu Menu
)

// Menu represents the horizontal Menu
type Menu struct {
	Items []Item `json:"MenuItems"`
}

// Item represents a menu item
type Item struct {
	InternalName    string `json:"internalName"`
	PublicName      string `json:"publicName"`
	Route           string `json:"route"`
	Description     string `json:"description"`
	IsVisibleInMain bool   `json:"isVisibleInMain"`
}

// SetConfig stores app config
func SetConfig(m Menu) {
	menu = m
}

// Config returns the config object
func Config() Menu {
	return menu
}
