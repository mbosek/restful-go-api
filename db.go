package main

import (
	"gopkg.in/mgo.v2"
)

func getDatabse() *mgo.Session {
	mongo, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	return mongo
}
