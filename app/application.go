package app

import (
	"github.com/gin-gonic/gin"
	"github.com/santiceron023/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("starting app")
	router.Run(":8080")

}
