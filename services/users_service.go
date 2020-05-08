package services

import (
	"github.com/santiceron023/bookstore_users-api/domain/users"
	"github.com/santiceron023/bookstore_users-api/utils/crypto_utils"
	"github.com/santiceron023/bookstore_users-api/utils/errors"
)

type usersService struct {
}

type usersServiceInterface interface {
	Search(string) (users.Users, *errors.RestError)
	DeleteUSer(int64) *errors.RestError
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUser(int64) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors.RestError)
}

var (
	UsersService usersServiceInterface = &usersService{}
)

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	//user.Email --> check NOT HERE
	//users.validate(&user) ---> NOT here

	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Password = crypto_utils.GetMd5(user.Password)
	user.Status = users.StatusActive
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) UpdateUser(partial bool, user users.User) (*users.User, *errors.RestError) {
	currentUSer, err := UsersService.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if partial {
		if user.FirstName != "" {
			currentUSer.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUSer.LastName = user.LastName
		}
		if user.Email != "" {
			currentUSer.Email = user.Email
		}
	} else {
		currentUSer.FirstName = user.FirstName
		currentUSer.LastName = user.LastName
		currentUSer.Email = user.Email
	}

	updateErr := currentUSer.Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return currentUSer, nil
}

func (s *usersService) DeleteUSer(userId int64) *errors.RestError {
	user := &users.User{Id: userId}
	deleteErr := user.Delete()
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}

func (s *usersService) Search(status string) (users.Users, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
