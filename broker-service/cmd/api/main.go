package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

// what is Addr?
// Addr is a string passed to net.Listen() function
// https://github.com/golang/go/blob/c7ea20195a3415668047eebdc488a4af1f629f04/src/net/http/server.go#L3247-L3255


// ln, err := net.Listen("tcp", ":8080")

// net.Listen returns net.TCPconn that has Read Write methods.