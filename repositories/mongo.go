package repositories
//TODO: check crud for mongo
// //
// // import (
// // 	"gopkg.in/mgo.v2"
// // )
// //
// // func getDatabse() *mgo.Session {
// // 	mongo, err := mgo.Dial("mongodb://localhost:27017")
// // 	if err != nil {
// // 		panic(err)
// // 	}
// //
// // 	return mongo
// // }
//
// package repository
//
// import(
//   _ "github.com/go-sql-driver/mysql"
//   "database/sql"
// )
//
// type dbConnection func() (db *mgo.Session)
//
// var Db dbConnection
//
// func connection() (db *mgo.Session) {
//
// 	dbDriver := "mongodb"
// 	dbAddress := "localhost"
// 	dbPort := "27017"
//
// 	db, err := mgo.Dial(dbDriver + "://"+dbAddress+":"+dbPort)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }
//
//
// // func init(){
// // 	Db = connection
// // }
//
// func get(){
//   db := connection()
// 	err := db.DB("rest").C("users").Find(nil).All(&users)
// 	if err != nil {
// 		return
// 	}
//
// }
//
// func create(){
//
// 	db := connection()
// 	err = db.DB("rest").C("users").Insert(user)
//
// 	if err != nil {
// 		if mgo.IsDup(err) {
// 			loggers.ErrorWithJSON(w, "User already exists", http.StatusBadRequest)
// 			return
// 		}
// 		loggers.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
// 		return
// 	}
// }
//
// func update(){
//   db := connection()
// 	err = uc.mongo.DB("rest").C("users").Update(bson.M{"name": userId}, &user)
// 	if err != nil {
// 			loggers.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
// 			return
// 	}
// }
//
// func delete(){
//   db := connection()
// 	err := uc.mongo.DB("rest").C("users").Remove(bson.M{"name": userId})
// 	if err != nil {
// 			switch err {
// 			default:
// 					loggers.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
// 					return
// 			case mgo.ErrNotFound:
// 					loggers.ErrorWithJSON(w, "Book not found", http.StatusNotFound)
// 					return
// 			}
// 	}
// }
