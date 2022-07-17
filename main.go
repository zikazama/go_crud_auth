package main

import (
	"go_crud_auth/config"
	middlewares "go_crud_auth/middlewares"
	routes_auth "go_crud_auth/routes/auth"
	routes_student "go_crud_auth/routes/student"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	router := gin.Default()

	router.GET("/", index)

	// route user
	router.POST("/login", routes_auth.LoginHandler)
	router.POST("/student", middlewares.Auth, routes_student.StoreStudent)
	router.GET("/student", middlewares.Auth, routes_student.ReadStudent)
	router.GET("/student/:id", middlewares.Auth, routes_student.ReadStudent)
	router.PUT("/student/:id", middlewares.Auth, routes_student.UpdateStudent)
	router.DELETE("/student/:id", middlewares.Auth, routes_student.DeleteStudent)

	router.Run()
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Halo ini API gin with gorm by Fauzi",
	})
}
