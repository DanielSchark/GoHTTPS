package main

import (
	"fmt"
	"log"
	"net/http"

	"./httpServer"
	"github.com/gorilla/mux"
)

var httpHandler httpServer.Handler

func redirectHTTPS(respWriter http.ResponseWriter, request *http.Request) {
	http.Redirect(
		respWriter,
		request,
		"https://127.0.0.1:10443/"+request.RequestURI,
		http.StatusMovedPermanently,
	)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", httpHandler.ServeHTTP)
	fmt.Printf("Now listening")
	go http.ListenAndServe(":80", http.HandlerFunc(redirectHTTPS))
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", router)
	if err != nil {
		log.Fatal(err)
	}
}
