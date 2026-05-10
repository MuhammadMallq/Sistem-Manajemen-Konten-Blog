package controllers

import (
	"net/http"

	"github.com/blog-cms/backend/config"
	"github.com/blog-cms/backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllComments(c *gin.Context) {
	var comments []models.Comment

	query := config.DB.Preload("Article")

	
	if articleID := c.Query("article_id"); articleID != "" {
		query = query.Where("article_id = ?", articleID)
	}

	result := query.Order("created_at DESC").Find(&comments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data komentar",
		"data":    comments,
	})
}


func GetCommentByID(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	result := config.DB.Preload("Article").First(&comment, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Komentar tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data komentar",
		"data":    comment,
	})
}

// CreateComment - Membuat komentar baru
func CreateComment(c *gin.Context) {
	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pastikan artikel ada
	var article models.Article
	if err := config.DB.First(&article, input.ArticleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	result := config.DB.Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Reload dengan relasi
	config.DB.Preload("Article").First(&input, input.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil membuat komentar baru",
		"data":    input,
	})
}

// UpdateComment - Memperbarui data komentar
func UpdateComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Komentar tidak ditemukan"})
		return
	}

	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&comment).Updates(models.Comment{
		CommenterName:  input.CommenterName,
		CommenterEmail: input.CommenterEmail,
		Content:        input.Content,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil memperbarui data komentar",
		"data":    comment,
	})
}

// DeleteComment - Menghapus komentar
func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Komentar tidak ditemukan"})
		return
	}

	config.DB.Delete(&comment)
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menghapus komentar",
	})
}

// GetCommentsByArticleID - Mengambil komentar berdasarkan Article ID
func GetCommentsByArticleID(c *gin.Context) {
	articleID := c.Param("id")
	var comments []models.Comment
	result := config.DB.Where("article_id = ?", articleID).Order("created_at DESC").Find(&comments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil komentar artikel",
		"data":    comments,
	})
}
