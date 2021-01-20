package handler

import (
	"gopkg.in/mgo.v2"
)

// Handler for db
type (
	Handler struct {
		DB *mgo.Session
	}
)

// Key
const (
	Key = "secret"
)
