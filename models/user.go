package models

// import (
	// "time"
// )

type (
	User struct {
		Id          int        `json:"id"`
		Name      	string     `json:"name"`
		Gender    	string     `json:"gender"`
		Age       	int        `json:"age"`
		// Created     time.Time  `json:"created_on"`
		// Updated     time.Time  `json:"updated_on"`
	}
)
