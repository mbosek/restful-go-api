package crud

import (
	"reflect"
	"strings"
	"net/http"

	"golang.org/x/net/context"
	"github.com/gedex/inflector"

	"github.com/mbosek/api/actions"
	"github.com/mbosek/api/router"
)

func Register(ctx context.Context, driver string, m interface{}) {
	if reflect.ValueOf(m).Kind() != reflect.Struct {
		panic("Model doesn't exist")
	}

	model := strings.ToLower(reflect.TypeOf(m).String())
	name := strings.Split(model, ".")
	pluralize := inflector.Pluralize(name[1])

	router.Add(name[1] + ".list", "GET", "/" + pluralize, handler(ctx, "list"))
	router.Add(name[1] + ".show", "GET", "/" + pluralize + "/{id}", handler(ctx, "show"))
	router.Add(name[1] + ".create", "POST", "/" + pluralize, handler(ctx, "create"))
	router.Add(name[1] + ".update", "PUT", "/" + pluralize + "/{id}", handler(ctx, "update"))
	router.Add(name[1] + ".delete", "DELETE", "/" + pluralize + "/{id}", handler(ctx, "delete"))
}

func handler(ctx context.Context, act string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		actions.Actions[act].Handle(ctx, w, r)
	}
}
