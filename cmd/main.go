package main

import (
	"log"
	"net/http"

	"github.com/CharlyF889/goSvelte/handler"
)

func main() {
	router := handler.NewRouter()
	router.ServeFiles("/static/*filepath", http.Dir("public"))
	router.Get("/", handler.Handler{H: handler.Index})
	router.Get("/shop", handler.Handler{H: handler.Shop})

	log.Fatal(http.ListenAndServe(":3000", router))
}
