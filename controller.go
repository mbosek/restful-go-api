package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"

	"github.com/mbosek/api/models"
	"github.com/mbosek/api/loggers"
)

type (
	UserController struct {
		mongo *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "RESTful API!")
}

func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	defer uc.mongo.Close()

	var users models.Users
	err := uc.mongo.DB("rest").C("users").Find(nil).All(&users)
	if err != nil {
		return
	}

	response, err := json.Marshal(users)

	loggers.ResponseWithJSON(w, response, http.StatusOK)

}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer uc.mongo.Close()

	var user models.User
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil {
		loggers.ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	err = uc.mongo.DB("rest").C("users").Insert(user)

	if err != nil {
		if mgo.IsDup(err) {
			loggers.ErrorWithJSON(w, "User already exists", http.StatusBadRequest)
			return
		}
		loggers.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		return
	}
	response, _ := json.Marshal(user)
	loggers.ResponseWithJSON(w, response, http.StatusCreated)
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if !bson.IsObjectIdHex(userId) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(userId)
	user := models.User{}

	if err := uc.mongo.DB("rest").C("users").FindId(oid).One(&user); err != nil {
		loggers.ErrorWithJSON(w, "User not exists", http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(user)
	loggers.ResponseWithJSON(w, response, http.StatusOK)
}

func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
	   loggers.ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
	   return
   	}

	// TODO: get by ID (mongo feature)
    err = uc.mongo.DB("rest").C("users").Update(bson.M{"name": userId}, &user)
    if err != nil {
        loggers.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
        return
    }

	response, _ := json.Marshal(user)
	loggers.ResponseWithJSON(w, response, http.StatusOK)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
        defer uc.mongo.Close()

		vars := mux.Vars(r)
		userId := vars["userId"]

		// TODO: delete by ID (mongo feature)
        err := uc.mongo.DB("rest").C("users").Remove(bson.M{"name": userId})
        if err != nil {
            switch err {
            default:
                loggers.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
                return
            case mgo.ErrNotFound:
                loggers.ErrorWithJSON(w, "Book not found", http.StatusNotFound)
                return
            }
        }

        w.WriteHeader(http.StatusNoContent)
}
