package main

import (
	"github.com/kzankpe/go-projects/url-shortening/routes"
)

func main() {
	// Initialize a router
	r := routes.InitRouter()
	routes.SetRoute(r)

	// Run the server on a port
	err := r.Run(":8090")
	if err != nil {
		panic(err)
	}
}
