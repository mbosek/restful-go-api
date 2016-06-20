package actions

import (
	"net/http"
	"fmt"
	"golang.org/x/net/context"
)

type createAction struct{}

func newCreateAction() Action {
	return &createAction{}
}

func (act *createAction) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request) {

}
