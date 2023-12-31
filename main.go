package main

import (
	"context"
	"net/http"
	"log"
	"os"
	"time"
	"Microservices/handlers"
	"os/signal"
)
func main(){
	
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	
	// create the handlers
	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)
	sm := http.NewServeMux()

	// create

	// sm.Handle("/",hh)
	// sm.Handle("/goodbye",gh)
	sm.Handle("/",ph)


	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func(){
		err := s.ListenAndServe()
		if err != nil{
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	s.Shutdown(tc)
}