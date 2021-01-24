package handler

import (
	"net/http"

	"github.com/ecila7290/instaGo/model"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// Signup method
func (h *Handler) Signup(c echo.Context) (err error) {
	u := &model.User{ID: bson.NewObjectId()}
	if err = c.Bind(u); err != nil {
		return
	}

	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("twitter").C("users").Insert(u); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, u)
}
