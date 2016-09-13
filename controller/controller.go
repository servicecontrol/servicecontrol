package controller

import (
	"servicecontrol.io/servicecontrol/controller/dashboard"
	"servicecontrol.io/servicecontrol/controller/static"
)

func LoadRoutes() {
	dashboard.Load()
	static.Load()
}
