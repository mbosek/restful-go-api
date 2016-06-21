package actions

import (
	"net/http"
	"strconv"

	"golang.org/x/net/context"
	"github.com/gorilla/mux"

	"github.com/mbosek/api/repositories"
)

func ShowAction() Action {
	return &showAction{}
}

func (act *showAction) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request) {
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
		r.Get(id, &model)
	}
}
