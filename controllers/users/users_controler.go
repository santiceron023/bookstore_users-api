package users

import (
	"github.com/gin-gonic/gin"
	"github.com/santiceron023/bookstore_users-api/domain/users"
	"github.com/santiceron023/bookstore_users-api/services"
	"github.com/santiceron023/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func getUserId(ginContext *gin.Context, key string) (int64, *errors.RestError) {
	userId, idUserErr := strconv.ParseInt(ginContext.Param(key), 10, 64)
	if idUserErr != nil {
		restErr := errors.NewBadRequestError("user ID should be a number")
		ginContext.JSON(restErr.Code, restErr)
		return 0, restErr
	}
	return userId, nil
}

func Create(ginContext *gin.Context) {
	var user users.User
	//bytes,err := ioutil.ReadAll(ginContext.Request.Body)
	//err := json.Unmarshal(bytes,&user);
	if err := ginContext.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("Invalid JSON")
		ginContext.JSON(restError.Code, restError)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		ginContext.JSON(saveErr.Code, saveErr)
		return
	}
	ginContext.JSON(http.StatusCreated, result)
}

func Get(ginContext *gin.Context) {
	//64 -> int64
	userId, idUserErr := getUserId(ginContext, "userId")
	if idUserErr != nil {
		ginContext.JSON(idUserErr.Code, idUserErr)
		return
	}

	result, restErr := services.GetUser(userId)
	if restErr != nil {
		ginContext.JSON(restErr.Code, restErr)
		return
	}
	ginContext.JSON(http.StatusOK, result.Marshal(ginContext.GetHeader("x-public") == "true"))

}

func Update(ginContext *gin.Context) {
	userId, idUserErr := getUserId(ginContext, "userId")
	if idUserErr != nil {
		ginContext.JSON(idUserErr.Code, idUserErr)
		return
	}

	var user users.User
	if err := ginContext.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("Invalid JSON")
		ginContext.JSON(restError.Code, restError)
		return
	}

	user.Id = userId

	partial := ginContext.Request.Method == http.MethodPatch

	result, updateErr := services.UpdateUser(partial, user)
	if updateErr != nil {
		ginContext.JSON(updateErr.Code, updateErr)
		return
	}
	ginContext.JSON(http.StatusOK, result.Marshal(ginContext.GetHeader("x-public") == "true"))

}

func Delete(ginContext *gin.Context) {
	userId, idUserErr := getUserId(ginContext, "userId")
	if idUserErr != nil {
		ginContext.JSON(idUserErr.Code, idUserErr)
		return
	}

	if err := services.DeleteUSer(userId); err != nil {
		ginContext.JSON(err.Code, err)
	}

	ginContext.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(ctx *gin.Context) {
	status := ctx.Query("status")
	users, findErr := services.FindByStatus(status)
	if findErr != nil {
		ctx.JSON(findErr.Code, findErr)
		return
	}
	users.Marshall(ctx.GetHeader("x-public") == "true")
	ctx.JSON(http.StatusOK, users)
}
