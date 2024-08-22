package handler

import (
	"net/http"
	"strconv"

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

// Get post

func GetPost(c *gin.Context) {
	postID := c.Param("id")
	id, err := strconv.Atoi(postID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var post model.Post

	if err := storage.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": "Invalid Post ID"})
		return
	}

	c.JSON(http.StatusOK, post)

}

// rep data needed to update post
type updateData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// update blog post

func UpdatePost(c *gin.Context) {
	var existingPost model.Post

	id := c.Param("id")

	if err := storage.DB.First(&existingPost, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	userID := c.MustGet("userID").(uint)

	if existingPost.AuthorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to edit this post"})
		return
	}

	var updateData updateData

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	existingPost.Title = updateData.Title
	existingPost.Content = updateData.Content

	// save the updated post
	if err := storage.DB.Save(&existingPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully!"})

}

// Delete Post

func DeletePost(c *gin.Context) {
	postID := c.Param("id")
	id, err := strconv.Atoi(postID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid post ID"})
		return
	}

	var post model.Post

	if err := storage.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": "Post not found"})
		return
	}

	if err := storage.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to delete Post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})

}
