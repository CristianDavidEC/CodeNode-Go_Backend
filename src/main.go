package main

import (
	"codenode/packages/src/router"
)

func main() {
	routerMux := router.Routing()
	server := NewServer(routerMux)
	server.Run()
}
