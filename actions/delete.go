package actions

import (
	"net/http"
	"fmt"
	"golang.org/x/net/context"
)

type deleteAction struct{}

func newDeleteAction() Action {
	return &deleteAction{}
}

func (act *deleteAction) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request) {
}
