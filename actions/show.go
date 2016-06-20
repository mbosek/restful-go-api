package actions

import (
	"net/http"
	"fmt"
	"golang.org/x/net/context"
)

type showAction struct{}

func newShowAction() Action {
	return &showAction{}
}

func (act *showAction) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request) {

}
