package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	path := flag.String("path", "/", "The path to serve from")
	port := flag.String("port", "3000", "The port to serve to on your localhost")
	flag.Parse()
	var address string = fmt.Sprintf(":%s", *port)

	http.Handle("/", http.FileServer(http.Dir(*path)))
	alert := fmt.Sprintf("Serving Assets from %s at %s", *path, *port)
	fmt.Println(alert)
	log.Fatal(http.ListenAndServe(address, nil))
}
