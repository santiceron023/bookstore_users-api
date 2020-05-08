package app

import (
	"github.com/santiceron023/bookstore_users-api/controllers/ping"
	"github.com/santiceron023/bookstore_users-api/controllers/users"
)

func mapUrls()  {

	router.GET("/ping", ping.Ping)
	router.GET("/users/:userId", users.Get)
	router.GET("/internal/users/search", users.Search)

	router.POST("/users", users.Create)
	router.PUT("/users/:userId", users.Update)
	router.PATCH("/users/:userId", users.Update)

	router.DELETE("/users/:userId", users.Delete)

}