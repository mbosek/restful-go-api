package actions

import (
	"net/http"
	"strconv"
	// "encoding/json"

	"golang.org/x/net/context"
	"github.com/gorilla/mux"

	"github.com/mbosek/api/repositories"
)

func DeleteAction() Action {
	return &deleteAction{}
}

func (act *deleteAction) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	if opt, ok := FromContext(ctx); ok {
		var (
			err error
			id int
		)

		vars := mux.Vars(req)
		id, err = strconv.Atoi(vars["id"]);
		if err != nil {
			panic(err)
		}

		model := opt.Model
		r := repositories.Repositories[opt.Repo]
		r.Delete(id, &model)
	}
}
