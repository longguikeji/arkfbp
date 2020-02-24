package main

import (
	"{{ .PackageName }}/flows/flow1"
	"{{ .PackageName }}/server"
)

// Routes ...
func Routes() []server.Route {
	var routes = []server.Route{
		server.Route{
			Name:    "flows.flow1",
			Pattern: "/hello",
			Handler: flow1.New(),
		},
	}

	return routes
}
