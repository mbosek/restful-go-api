package router

import (
	"net/http"
)

type (
	route struct {
		Method  string
		Pattern string
		Handler http.HandlerFunc
	}
	pool map[string]*route
)

var (
	routes = make(pool)
)

func Add(name, method, pattern string, handler http.HandlerFunc) {
	routes[name] = &route{method, pattern, handler}
}
