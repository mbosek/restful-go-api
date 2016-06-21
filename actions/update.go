package actions

import (
	"net/http"
	"strconv"
	"encoding/json"
	"log"

	"golang.org/x/net/context"
	"github.com/gorilla/mux"

	"github.com/mbosek/api/loggers"
	"github.com/mbosek/api/repositories"
)

func UpdateAction() Action {
	return &updateAction{}
}

func (act *updateAction) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request) {
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
		r.Update(id, &model)

		response, err := json.Marshal(&model)
        if err != nil {
            log.Fatal(err)
        }
		loggers.ResponseWithJSON(w, response, http.StatusOK)
	}
}
