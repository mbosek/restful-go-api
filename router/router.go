package router

import "github.com/gorilla/mux"
import "net/http"

type (
	routerFactory func() *mux.Router
	route struct {
		Method  string
		Pattern string
		Handler http.HandlerFunc
	}
	pool map[string]*route
)

var routes = make(pool)

var New routerFactory

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for name, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(name).
			Handler(route.Handler)
	}

	return router
}

func Add(name, method, pattern string, handler http.HandlerFunc) {
	routes[name] = &route{method, pattern, handler}
}

func init() {
	New = newRouter
}
