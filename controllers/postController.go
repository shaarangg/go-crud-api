package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shaarangg/go-rest-api/initializers"
	"github.com/shaarangg/go-rest-api/models"
)

func PostCreate(c *gin.Context) {
	//  Get Data of the body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.POST{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the data
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsAll(c *gin.Context) {
	var posts []models.POST
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostById(c *gin.Context) {
	var post models.POST

	postId := c.Param("id")

	initializers.DB.Find(&post, postId)
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	var post models.POST

	postId := c.Param("id")

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	initializers.DB.First(&post, postId)

	initializers.DB.Model(&post).Updates(models.POST{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	var post models.POST

	postId := c.Param("id")

	initializers.DB.Delete(&post, postId)

	c.JSON(200, gin.H{
		"message": "success",
	})
}
