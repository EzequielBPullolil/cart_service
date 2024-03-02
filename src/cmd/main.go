package main

import (
	"log"

	"github.com/EzequielBPullolil/cart_service/src"
)

func main() {
	app := src.CreateApp()

	log.Fatal(app.Run(":8030"))
}
