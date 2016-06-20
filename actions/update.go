package actions

import (
	"net/http"
	"fmt"
	"golang.org/x/net/context"
)

type updateAction struct{}

func newUpdateAction() Action {
	return &updateAction{}
}

func (act *updateAction) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request) {

}
