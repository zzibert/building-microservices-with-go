package main

import (
	"log"
	"net/http"
	"os"

	"github.com/zzibert/building-microservices-with-go/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// Listen for connections on all ip addresses (0.0.0.0)
	// port 9090
	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", sm)
	log.Fatal(err)
}
