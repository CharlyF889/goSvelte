package main

import (
	"log"
	"net/http"

	"github.com/CharlyF889/goSvelte/handler"
)

func main() {
	router := handler.NewRouter()
	router.Get("/", handler.Handler{H: handler.Index})
	log.Fatal(http.ListenAndServe(":3000", router))
}
