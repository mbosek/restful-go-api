package models

import (
	"time"
)

type (
	Book struct {
		Id        int        `json:"id"`
		Author    string     `json:"description"`
		Created   time.Time  `json:"created_on"`
		Updated   time.Time  `json:"updated_on"`
	}
)
