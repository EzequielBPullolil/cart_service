package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/EzequielBPullolil/cart_service/src"
)

func main() {
	var host string
	var port int

	flag.StringVar(&host, "host", "localhost", "Host to listen on")
	flag.IntVar(&port, "port", 8080, "Port to listen on")

	flag.Parse()
	app := src.CreateApp()

	log.Fatal(app.Run(fmt.Sprintf("%s:%d", host, port)))
}
