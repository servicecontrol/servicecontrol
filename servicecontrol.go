package main

import (
"servicecontrol.io/servicecontrol/boot"
"servicecontrol.io/servicecontrol/lib/router"
"servicecontrol.io/servicecontrol/lib/server"
)

func main() {
	info := boot.LoadConfig("env.json")

	boot.RegisterServices(info)

	handler := boot.SetUpMiddleware(router.Instance())

	server,Run (
		handler,
		handler,
		info.Server,
	)
}
