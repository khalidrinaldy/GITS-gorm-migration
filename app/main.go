package main

import "migration/routes"

func main() {
	e := routes.InitRoute()

	e.Start(":4000")
}
