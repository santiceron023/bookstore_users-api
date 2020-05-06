package users

import (
	"github.com/gin-gonic/gin"
	"github.com/santiceron023/bookstore_users-api/domain/users"
	"github.com/santiceron023/bookstore_users-api/services"
	"github.com/santiceron023/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	//bytes,err := ioutil.ReadAll(c.Request.Body)
	//err := json.Unmarshal(bytes,&user);
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("Invalid JSON")
		c.JSON(restError.Code, restError)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Code,saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	//64 -> int64
	userId, userErr := strconv.ParseInt(c.Param("userId"),10,64)
	if userErr != nil{
		err := errors.NewBadRequestError("user ID should be a number")
		c.JSON(err.Code,err)
		return
	}

	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}
