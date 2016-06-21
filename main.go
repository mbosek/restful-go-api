package main

import (
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/mbosek/api/router"
	"github.com/mbosek/api/crud"
	"github.com/mbosek/api/models"
	"github.com/mbosek/api/repositories"
)

func main() {

	var ctx context.Context
	ctx, _ = context.WithCancel(context.Background())

	repositories.Register("mysql", repositories.NewRepository("go"))
	modelCtrl := crud.New(ctx, "mysql", models.User{})
	modelCtrl.AddActions()

	log.Fatal(http.ListenAndServe(":8080", router.New()))
}
