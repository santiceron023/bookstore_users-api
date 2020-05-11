package users

import (
	"fmt"
	"github.com/santiceron023/FUENTE/bookstore_utils-go/rest_errors"
	"github.com/santiceron023/bookstore_users-api/datasources/mysql/users_db"
	"github.com/santiceron023/bookstore_users-api/logger"
	"github.com/santiceron023/bookstore_users-api/utils/date_utils"
	"github.com/santiceron023/bookstore_users-api/utils/errors"
	"github.com/santiceron023/bookstore_users-api/utils/mysql_utils"
	"strings"
)

const (
	QueryInsertUser             = "INSERT INTO users (first_name,last_name,email,date_created,status,password) VALUES (?,?,?,?,?,?)"
	QueryGetUserId              = "SELECT id,first_name,last_name,date_created,email,status,password FROM users WHERE id=?"
	QueryUpdateUser             = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?"
	QueryDeleteUser             = "DELETE FROM users WHERE id=?"
	QueryFindByStatus           = "SELECT id,first_name,last_name,date_created,email,status FROM users WHERE status=?"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)

func (user *User) Save() *errors.RestError {

	statement, err := users_db.Client.Prepare(QueryInsertUser)
	if err != nil {
		logger.Error("error when parsing user SQL INSERT", err)
		return errors.NewInternalServerError("DataBase Error")
	}
	defer statement.Close()

	user.DateCreated = date_utils.GetNowdbDBString()

	insertRes, saveErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when executing user SQL INSERT", err)
		return errors.NewInternalServerError("DataBase Error")
		//return mysql_utils.ParseError(saveErr)
	}
	userId, idErr := insertRes.LastInsertId()
	if idErr != nil {
		logger.Error("error when getting user last id SQL INSERT", err)
		return errors.NewInternalServerError("DataBase Error")
		//return mysql_utils.ParseError(idErr)
	}
	user.Id = userId
	return nil
}

func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(QueryGetUserId)
	if err != nil {
		logger.Error("error when parsing user SQL GET by ID", err)
		return errors.NewInternalServerError("DataBase Error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getError := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated, &user.Email, &user.Status, &user.Password); getError != nil {
		logger.Error("error when executing user SQL GET by ID", err)
		return errors.NewInternalServerError("DataBase Error")
	}
	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, queryErr := users_db.Client.Prepare(QueryUpdateUser)
	if queryErr != nil {
		return errors.NewInternalServerError(queryErr.Error())
	}
	defer stmt.Close()

	_, excErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if excErr != nil {
		return mysql_utils.ParseError(excErr)
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, queryErr := users_db.Client.Prepare(QueryDeleteUser)
	if queryErr != nil {
		logger.Error("error when parsing user SQL DELETE", queryErr)
		return errors.NewInternalServerError("DataBase Error")
	}
	defer stmt.Close()

	_, excErr := stmt.Exec(user.Id)
	if excErr != nil {
		logger.Error("error when executing user SQL DELETE", queryErr)
		return errors.NewInternalServerError("DataBase Error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {
	smt, queryErr := users_db.Client.Prepare(QueryFindByStatus)
	if queryErr != nil {
		logger.Error("error when parsing user SQL SEARCH", queryErr)
		return nil, errors.NewInternalServerError("DataBase Error")
	}
	defer smt.Close()

	resultRows, err := smt.Query(status)
	if err != nil {
		logger.Error("error when executing user SQL SEARCH", queryErr)
		return nil, errors.NewInternalServerError("DataBase Error")
	}
	defer resultRows.Close()

	results := make([]User, 0)
	for resultRows.Next() {
		var UserRes User
		if err := resultRows.Scan(&UserRes.Id, &UserRes.FirstName, &UserRes.LastName, &UserRes.DateCreated, &UserRes.Email, &UserRes.Status); err != nil {
			logger.Error("error when scanning row to user struct", queryErr)
			return nil, errors.NewInternalServerError("DataBase Error")
		}
		results = append(results, UserRes)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users with status %s", status))
	}

	return results, nil
}

func (user *User) FindByEmailAndPassword() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return rest_errors.NewInternalServerError("error when tying to find user", err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return rest_errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password", getErr)
		return rest_errors.NewInternalServerError("error when tying to find user", getErr)
	}
	return nil
}
