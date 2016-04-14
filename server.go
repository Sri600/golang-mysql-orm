package main

import (
	// Standard library packages

	// Third party packages
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	//"os"
	"./controllers"
)


func main() {

	// Get DBController instance
	dc := controllers.DBController{}
	dc.InitDB()
    dc.InitSchema()

	// Get a TodolistController instance
	todouser := controllers.TodolistControllerUser{}
	todouser.SetDB(dc.GetDB())

	todolist := controllers.TodolistController{}
	todolist.SetDB(dc.GetDB())




	// Get a todolist resource
	router := gin.Default()

	router.GET("/users", todouser.ListUsers)
    router.GET("/users/:id", todouser.GetUsers)
    router.DELETE("/users/:id", todouser.DeleteUsers)
    router.POST("/users", todouser.CreateUsers)
    router.PUT("/users/:id", todouser.UpdateUsers)

    router.GET("/todolist", todolist.ListTodolist)
    router.GET("/todolist/:id", todolist.GetTodolist)
    router.DELETE("/todolist/:id", todolist.DeleteTodolist)
    router.POST("/todolist", todolist.CreateTodolist)
    router.PUT("/todolist/:id", todolist.UpdateTodolist)

    router.Run(":8000")
}
