package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hazaloolu/blog-api/internal/model"
	"github.com/hazaloolu/blog-api/internal/storage"
)

// create blog post

func CreatePost(c *gin.Context) {
	var post model.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)
	post.AuthorID = userID

	if err := storage.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully!"})
}

// rep data needed to update post
type updateData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// update blog post

func UpdatePost(c *gin.Context) {
	var existingPost model.Post

	// get post id from url
	id := c.Param("id")

	// get the post using the id

	if err := storage.DB.First(&existingPost, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// check if the authenticated user is the author of the post
	userID := c.MustGet("userID").(uint)

	if existingPost.AuthorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to edit this post"})
		return
	}

	// Bind the JSON payload to an update post struct

	var updateData updateData

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	// update the post with the new title and content
	existingPost.Title = updateData.Title
	existingPost.Content = updateData.Content

	// save the updated post
	if err := storage.DB.Save(&existingPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully!"})

}
