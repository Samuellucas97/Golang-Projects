package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Give a behavior to defaultServiceMux (default HTTP request handlers)
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println("Hello :)")

		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "The url is malformed", http.StatusBadRequest)
			return
		}

		log.Printf("Data: %s", data)

		fmt.Fprintf(res, "Hi %s! How've you been?\n", data)
	})

	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("Bye ;) S2")
	})

	// You could also use just ":9090"
	http.ListenAndServe("127.0.0.1:9090", nil)
}

