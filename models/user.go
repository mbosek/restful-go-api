package models

import (
 	"time"
)

type (
	User struct {
		Name      string     `json:"name" bson:"name"`
		Gender    string     `json:"gender" bson:"gender"`
		Age       int        `json:"age" bson:"age"`
		Created   time.Time  `json:"created" bson:"created"`
	}
)

type Users []User
