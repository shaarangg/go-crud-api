package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shaarangg/go-rest-api/controllers"
	"github.com/shaarangg/go-rest-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostsAll)
	r.GET("/posts/:id", controllers.PostById)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.Run()
}
