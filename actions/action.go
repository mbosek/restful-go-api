package actions

import (
	"net/http"
	"golang.org/x/net/context"
)

type (
	Action interface {
		Handle(ctx context.Context, w http.ResponseWriter, req *http.Request)
	}
	Options struct {
		Model  interface{}
		Repo string
	}
	serviceKey struct{}
	actions map[string]Action
)

var Actions = make(actions)

func FromContext(ctx context.Context) (Options, bool) {
	o, ok := ctx.Value(serviceKey{}).(Options)
	return o, ok
}

func init() {
	Actions["list"] = newListAction()
	Actions["show"] = newShowAction()
	Actions["create"] = newCreateAction()
	Actions["update"] = newUpdateAction()
	Actions["delete"] = newDeleteAction()
}
