package main

import "codenode.com/packages/src/router"

func main() {

	routerMux := router.Routing()
	server := NewServer(routerMux)
	server.Run()
}
