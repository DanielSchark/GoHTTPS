package httpServer

import (
	"net/http"
)

type Handler struct{}

func (handler Handler) ServeHTTP(respWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {

	case "GET":
		respWriter.Header().Set("Content-Type", "text/plain")
		respWriter.Write([]byte(
			"Here you go with your GET request!",
		))

	case "POST":
		respWriter.Header().Set("Content-Type", "text/plain")
		respWriter.Write([]byte(
			"Thanks for the post!",
		))

	}
}
