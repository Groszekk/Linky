package main

import (
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
	linkControlers := ServiceContainer().InjectController()

	r.HandleFunc("/short", linkControlers.ShortLink).Methods("POST")

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
