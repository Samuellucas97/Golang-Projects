package main

import (
	"awesomeProject/fivethClass/product-api/handlers"
	"context"
	//"github.com/nicholasjackson/env"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//var binAddress = env.String("BIND_ADDRESS", false, "localhost:9090", "Bind address for the server")

func main() {

	l := log.New(os.Stdout, "my-api", log.LstdFlags)

	productHandler := handlers.NewProducts(l)

	myServeMux := mux.NewRouter()

	getRouter := myServeMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productHandler.GetProducts)

	postRouter := myServeMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProduct)

	putRouter := myServeMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProduct)

	myServer := &http.Server{
		Addr: ":9090",
		Handler: myServeMux,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 3*time.Second,
	}

	// Using goroutine here
	go func() {
		err := myServer.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <- sigChannel
	l.Println("Received terminate signal, graceful shutdown %s",  sig)

	timeContext, _ := context.WithTimeout(context.Background(), 60*time.Second)
	myServer.Shutdown(timeContext) // It's necessary to system reliability, creating grace period
}

