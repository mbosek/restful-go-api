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
		Model interface{}
		Repo  string
	}
	listAction struct{}
	showAction struct{}
	createAction struct{}
	updateAction struct{}
	deleteAction struct{}

	serviceKey struct{}
	actions map[string]Action
)


var Actions = make(actions)

func NewContext(ctx context.Context, o Options) context.Context {
	return context.WithValue(ctx, serviceKey{}, o)
}

func FromContext(ctx context.Context) (Options, bool) {
	o, ok := ctx.Value(serviceKey{}).(Options)
	return o, ok
}

func init() {
	Actions["list"] = ListAction()
	Actions["show"] = ShowAction()
	Actions["create"] = CreateAction()
	Actions["update"] = UpdateAction()
	Actions["delete"] = DeleteAction()
}
