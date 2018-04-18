package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"./resolver"
	"./schema"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var (
	host              string
	port              string
	readHeaderTimeout = 1 * time.Second
	writeTimeout      = 10 * time.Second
	idleTimeout       = 90 * time.Second
	graph             *graphql.Schema
)

func main() {
	flag.StringVar(
		&host,
		"host", "localhost",
		"Host address like 127.0.0.1",
	)
	flag.StringVar(
		&port,
		"port", "80",
		"Port like 443 for HTTPS",
	)
	flag.Parse()
	graph = graphql.MustParseSchema(
		schema.Schema,
		&resolver.Resolver{},
	)

	// GraphQL handler
	http.Handle("/graph", &relay.Handler{Schema: graph})

	// listenOnHTTPS := http.Server{
	// 	Addr:              host + ":" + port,
	// 	ReadHeaderTimeout: readHeaderTimeout,
	// 	WriteTimeout:      writeTimeout,
	// 	IdleTimeout:       idleTimeout,
	// }

	fmt.Printf(
		"---- Now listening on %v:%v ----\n",
		// "Unsecure connections(HTTP) "+
		// "will be redirected to secure(HTTPS)\n",
		host,
		port,
	)
	// Redirect to HTTPS when querying throught HTTP
	// go http.ListenAndServe(
	// 	host+":80",
	// 	http.HandlerFunc(http.HandlerFunc(
	// 		func(
	// 			respWriter http.ResponseWriter,
	// 			request *http.Request,
	// 		) {
	// 			http.Redirect(
	// 				respWriter,
	// 				request,
	// 				"https://"+host+
	// 					":443"+request.RequestURI,
	// 				http.StatusTemporaryRedirect,
	// 			)
	// 		},
	// 	)),
	// )

	// Listen and serve over TLS
	// err := listenOnHTTPS.ListenAndServeTLS("cert.pem", "key.pem")
	err := http.ListenAndServe(host+":"+port, nil)
	fmt.Println("Shuting down server")
	if err != nil {
		log.Fatal(err)
	}
}
