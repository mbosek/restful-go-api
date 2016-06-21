package crud

import (
	// "fmt"
	"reflect"
	"strings"
	"net/http"

	"golang.org/x/net/context"
	"github.com/gedex/inflector"

	"github.com/mbosek/api/actions"
	"github.com/mbosek/api/router"
)

type (
	Controller interface {
		AddActions(...string)
	}

	controller struct {
		ctx   context.Context
		model interface{}
		repo  string
	}
)

var New = newController

func (ctrl *controller) AddActions(acts ...string) {
	if reflect.ValueOf(ctrl.model).Kind() != reflect.Struct {
		panic("Model doesn't exist")
	}

	model := strings.ToLower(reflect.TypeOf(ctrl.model).String())
	name := strings.Split(model, ".")
	pluralize := inflector.Pluralize(name[1])

	router.Add(name[1] + ".list", "GET", "/" + pluralize, handler(ctrl.ctx, "list"))
	router.Add(name[1] + ".show", "GET", "/" + pluralize + "/{id}", handler(ctrl.ctx, "show"))
	router.Add(name[1] + ".create", "POST", "/" + pluralize, handler(ctrl.ctx, "create"))
	router.Add(name[1] + ".update", "PUT", "/" + pluralize + "/{id}", handler(ctrl.ctx, "update"))
	router.Add(name[1] + ".delete", "DELETE", "/" + pluralize + "/{id}", handler(ctrl.ctx, "delete"))
}

func handler(ctx context.Context, act string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		actions.Actions[act].Handle(ctx, w, r)
	}
}

func newController(ctx context.Context, repo string, m interface{}) Controller {
	ctx = actions.NewContext(ctx, actions.Options{m, repo})
	return &controller{ctx, m, repo}
}
