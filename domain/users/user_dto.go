package users

import (
	"github.com/santiceron023/bookstore_users-api/utils/errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json: "id"`
	FirstName   string `json: "firstName"`
	LastName    string `json: "lastName"`
	Email       string `json: "email"`
	DateCreated string `json: "dateCreated"`
	Status string `json: "status"`
	Password string `json: "password"` //internal field

}

type Users []User

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}

//func (user *User) validatePass
