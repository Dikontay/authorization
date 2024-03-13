package main

import (
	"auth/controllers"
	"auth/pkg"
	"auth/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	pkg.LoadEnvVariables()
	pkg.ConnectToDB()
	pkg.SyncDatabase()
}
func main() {

	router := gin.Default()
	authRoutes := router.Group("/auth/user")
	// registration route
	authRoutes.POST("/register", controllers.Signup)
	// login route
	authRoutes.POST("/login", controllers.Login)

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(utils.JwtAuth)
	adminRoutes.GET("/users", controllers.GetUsers)
	adminRoutes.GET("/user/:id", controllers.GetUser)
	adminRoutes.PUT("/user/:id", controllers.UpdateUser)
	adminRoutes.POST("/user/role", controllers.CreateRole)
	adminRoutes.GET("/user/roles", controllers.GetRoles)
	adminRoutes.PUT("/user/role/:id", controllers.UpdateRole)

	err := router.Run(":8000")
	if err != nil {
		return
	}
	fmt.Println("Server running on port 8000")

}
