package actions

import (
	"net/http"
	"reflect"
	
	"golang.org/x/net/context"

	"github.com/mbosek/api/repositories"

)

func ListAction() Action {
	return &listAction{}
}

func (act *listAction) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	if opt, ok := FromContext(ctx); ok {
		var (
			// err error
			// body []byte
		)
		myType := reflect.TypeOf(opt.Model)
		models := reflect.Zero(reflect.SliceOf(myType)).Interface()
		r := repositories.Repositories[opt.Repo]
		r.GetAll(&models)
	}
}
