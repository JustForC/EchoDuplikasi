package main

import (
	"Kynesia/routes"
)

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":4562"))
}
