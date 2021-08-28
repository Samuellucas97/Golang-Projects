package main

import (
	"awesomeProject/thirdClass/product-api/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	l := log.New(os.Stdout, "my-api", log.LstdFlags)

	myServeMux := http.NewServeMux()
	productHandler := handlers.NewProducts(l)
	myServeMux.Handle("/", productHandler)

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

