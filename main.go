package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"./httpHandler"
	"github.com/gorilla/mux"
)

var (
	host           string
	port           string
	graphqlHandler httpHandler.GraphqlHandler
)

func redirectHTTPS(
	respWriter http.ResponseWriter,
	request *http.Request,
) {
	// fmt.Println("Here in REDIRECT body:", request.Body)
	http.Redirect(
		respWriter,
		request,
		"https://127.0.0.1:443"+request.RequestURI,
		http.StatusTemporaryRedirect,
	)
}

func main() {
	flag.StringVar(
		&host, "host",
		"0.0.0.0", "Host address like 127.0.0.1",
	)
	flag.StringVar(
		&port, "port",
		"443", "Port like 443 for HTTPS",
	)
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", graphqlHandler.ServeHTTP)

	fmt.Printf("Now listening on %v:%v\n", host, port)

	// Redirect to HTTPS when querying throught HTTP
	go http.ListenAndServe(":80", http.HandlerFunc(redirectHTTPS))

	// Listen and serve TLS - Secure
	err := http.ListenAndServeTLS(
		host+":"+port,
		"cert.pem", "key.pem",
		router,
	)
	if err != nil {
		log.Fatal(err)
	}
}
