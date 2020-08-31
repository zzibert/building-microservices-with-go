package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "00000sps", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", data)

	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", nil)
}
