package users

import (
	"fmt"
	"github.com/santiceron023/bookstore_users-api/utils/errors"
)
var(
	usersDb = make(map[int64]*User)
)

func (user *User)Save() *errors.RestError{
	if current := usersDb[user.Id]; current != nil {
		if current.Email == user.Email{
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered",user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user with id %d already exists",user.Id))
	}
	usersDb[user.Id] = user
	return nil
}

func (user *User)Get() *errors.RestError{
	result := usersDb[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %v not found", user.Id))
	}
	//row
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.LastName
	return nil
}