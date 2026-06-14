package routes

import (
	"github-analyzer/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/users", handlers.CreateUser)

	r.GET("/users", handlers.GetAllUsers)

	r.GET("/users/:id", handlers.GetUserByID)

	r.PUT("/users/:id", handlers.UpdateUser)

	r.DELETE("/users/:id", handlers.DeleteUser)

	r.POST("/login", handlers.Login)

	r.GET("/github/:username", handlers.GetGithubUser)
}
