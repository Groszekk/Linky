package main

import (
	"Linky/ShorterService/middlewares"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Router interface {
	Init() *mux.Router
}

type router struct{}

var (
	r     *router
	rOnce sync.Once
)

func (router *router) Init() *mux.Router {
	r := mux.NewRouter()
	linkControllers := ServiceContainer().InjectController()

	r.Handle("/short", middlewares.Validation(http.HandlerFunc(linkControllers.ShortLink))).Methods("POST")
	r.HandleFunc("/s/{short}", linkControllers.ShortRedirect).Methods("GET")

	return r
}

func HttpRouter() Router {
	if r == nil {
		rOnce.Do(func() {
			r = &router{}
		})
	}

	return r
}
