package main

import (
	"github-analyzer/internal/database"
	"github-analyzer/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectMongo()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Static("/css", "./web/css")
	r.Static("/js", "./web/js")
	r.StaticFile("/", "./web/index.html")
	r.StaticFile("/users-page", "./web/users.html")

	r.Run(":8080")
}
