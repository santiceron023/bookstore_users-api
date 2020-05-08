package users

import (
	"fmt"
	"github.com/santiceron023/bookstore_users-api/datasources/mysql/users_db"
	"github.com/santiceron023/bookstore_users-api/utils/date_utils"
	"github.com/santiceron023/bookstore_users-api/utils/errors"
	"github.com/santiceron023/bookstore_users-api/utils/mysql_utils"
)

const (
	QueryInsertUser = "INSERT INTO users(first_name,last_name,email,date_created,status,password) VALUES (?,?,?,?,?,?)"
	QueryGetUserId  = "SELECT id,first_name,last_name,date_created,email,status,password FROM users WHERE id=?"
	QueryUpdateUser  = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?"
	QueryDeleteteUser  = "DELETE FROM users WHERE id=?"
	QueryFindByStatus  = "SELECT id,first_name,last_name,date_created,email,status FROM users WHERE status=?"
)

func (user *User) Save() *errors.RestError {

	statement, err := users_db.Client.Prepare(QueryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer statement.Close()

	user.DateCreated = date_utils.GetNowdbDBString()

	insertRes, saveErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated,user.Status,user.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, idErr := insertRes.LastInsertId()
	if idErr != nil {
		return mysql_utils.ParseError(idErr)
	}
	user.Id = userId
	return nil
}

func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(QueryGetUserId)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getError := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated, &user.Email,&user.Status,&user.Password); getError != nil {
		return mysql_utils.ParseError(getError)
	}
	return nil
}

func (user *User) Update() *errors.RestError  {
	stmt,queryErr := users_db.Client.Prepare(QueryUpdateUser)
	if queryErr != nil{
		return errors.NewInternalServerError(queryErr.Error())
	}
	defer stmt.Close()

	_, excErr := stmt.Exec(user.FirstName,user.LastName,user.Email,user.Id)
	if excErr != nil{
		return mysql_utils.ParseError(excErr)
	}
	return nil
}

func (user *User) Delete() *errors.RestError  {
	stmt,queryErr := users_db.Client.Prepare(QueryDeleteteUser)
	if queryErr != nil{
		return errors.NewInternalServerError(queryErr.Error())
	}
	defer stmt.Close()

	_, excErr := stmt.Exec(user.Id)
	if excErr != nil{
		return mysql_utils.ParseError(excErr)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User,*errors.RestError){
	smt,err := users_db.Client.Prepare(QueryFindByStatus)
	if err != nil{
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer smt.Close()

	resultRows,err := smt.Query(status)
	if err != nil{
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer resultRows.Close()

	results := make([]User,0)
	for resultRows.Next(){
		var UserRes User
		if err := resultRows.Scan(&UserRes.Id,&UserRes.FirstName,&UserRes.LastName,&UserRes.DateCreated,&UserRes.Email,&UserRes.Status);err != nil{
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results,UserRes)
	}

	if len(results) ==0{
		return nil,errors.NewNotFoundError(fmt.Sprintf("no users with status %s",status))
	}

	return results,nil



}
