package users

import "encoding/json"

type PublicUser struct{
	Id          int64  `json: "id"`
	DateCreated string `json: "dateCreated"`
	Status string `json: "status"`
	}

type PrivateUser struct{
	Id          int64  `json: "id"`
	FirstName   string `json: "firstName"`
	LastName    string `json: "lastName"`
	Email       string `json: "email"`
	DateCreated string `json: "dateCreated"`
	Status string `json: "status"`
	//Password string `json: "password"` //internal field
	}

	func (users *Users) Marshall(isPublic bool) []interface{} {
		result := make([]interface{},len(*users))
		for index,user := range *users{
			result[index] = user.Marshal(isPublic)
		}
		return result
	}

		func (user *User) Marshal(isPublic bool) interface{}{
		if isPublic{
			// is assigned like this if JSON != id or names
			return PublicUser{
				Id: user.Id,
				DateCreated: user.DateCreated,
				Status: user.Status,
			}
		}

		userJson,_ := json.Marshal(user)
		var privateUser PrivateUser
		json.Unmarshal(userJson,&privateUser)
		return privateUser

		}

