package main

import (
	"servicecontrol.io/servicecontrol/boot"
	"servicecontrol.io/servicecontrol/lib/router"
	"servicecontrol.io/servicecontrol/lib/server"
)

// main loads the configuration file, registers the services, applies the
// middleware to the router, and then starts the HTTP and HTTPS listeners.
func main() {
	// Load the configuration file
	info := boot.LoadConfig("env.json")

	// Register the services
	boot.RegisterServices(info)

	// Retrieve the middleware
	handler := boot.SetUpMiddleware(router.Instance())

	// Start the HTTP and HTTPS listeners
	server.Run(
		handler,     // HTTP handler
		handler,     // HTTPS handler
		info.Server, // Server settings
	)
}
