package mysql_utils

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/santiceron023/bookstore_users-api/utils/errors"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewBadRequestError(fmt.Sprintf("no record matching id"))
		}
		return errors.NewInternalServerError("Error parsing DataBase response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("Invalidad Data, no repeated something"))
	}

	return errors.NewInternalServerError("error processing request")
}
