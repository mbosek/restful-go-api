package main

import (
	"github.com/mbosek/api/models"
)

var db = getDatabse()
var uc = NewUserController(db)

var routes = models.Routes{
	models.Route{
		"Index",
		"GET",
		"/",
		uc.Index,
	},
	models.Route{
		"GetUsers",
		"GET",
		"/users",
		uc.GetUsers,
	},
	models.Route{
		"CreateUser",
		"POST",
		"/users",
		uc.CreateUser,
	},
	models.Route{
		"GetUser",
		"GET",
		"/users/{userId}",
		uc.GetUser,
	},
	models.Route{
		"UpdateUser",
		"PUT",
		"/users/{userId}",
		uc.UpdateUser,
	},
	models.Route{
		"DeleteUser",
		"DELETE",
		"/users/{userId}",
		uc.DeleteUser,
	},
}
