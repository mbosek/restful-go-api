package actions

import (
	"net/http"
	"fmt"
	"golang.org/x/net/context"
)

type listAction struct{}

func newListAction() Action {
	return &listAction{}
}

func (act *listAction) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request) {

}
