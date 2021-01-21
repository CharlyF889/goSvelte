package handler

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

type key int

const psKey key = 0

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type Router struct {
	*httprouter.Router
}

func (r *Router) Get(path string, handler http.Handler) {
	r.GET(path, wrapHandler(handler))
}
func (r *Router) Post(path string, handler http.Handler) {
	r.POST(path, wrapHandler(handler))
}
func (r *Router) Put(path string, handler http.Handler) {
	r.PUT(path, wrapHandler(handler))
}
func (r *Router) Delete(path string, handler http.Handler) {
	r.DELETE(path, wrapHandler(handler))
}

func NewRouter() *Router {
	return &Router{httprouter.New()}
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctxWithParams := context.WithValue(r.Context(), psKey, ps)
		rWithPS := r.WithContext(ctxWithParams)
		h.ServeHTTP(w, rWithPS)
	}
}

type Handle func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error

type Handler struct {
	H func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ps := r.Context().Value(psKey).(httprouter.Params)
	err := h.H(w, r, ps)
	if err != nil {
		switch e := err.(type) {
		case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	fmt.Println("Connection")

	base := filepath.Join("templates", "base.html")
	index := filepath.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(base, index)
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "base", nil)

	return nil
}
func Shop(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	fmt.Println("Connection")

	base := filepath.Join("templates", "base.html")
	shop := filepath.Join("templates", "shop.html")
	tmpl, err := template.ParseFiles(base, shop)
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "base", nil)

	return nil
}
