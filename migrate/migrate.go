package main

import (
	"github.com/shaarangg/go-rest-api/initializers"
	"github.com/shaarangg/go-rest-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.POST{})
}
