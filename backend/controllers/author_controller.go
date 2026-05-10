package controllers

import (
	"net/http"

	"github.com/blog-cms/backend/config"
	"github.com/blog-cms/backend/models"
	"github.com/gin-gonic/gin"
)

// GetAllAuthors - Mengambil semua data penulis
func GetAllAuthors(c *gin.Context) {
	var authors []models.Author
	result := config.DB.Preload("Articles").Find(&authors)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data penulis",
		"data":    authors,
	})
}

// GetAuthorByID - Mengambil data penulis berdasarkan ID
func GetAuthorByID(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	result := config.DB.Preload("Articles").First(&author, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Penulis tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data penulis",
		"data":    author,
	})
}

// CreateAuthor - Membuat penulis baru
func CreateAuthor(c *gin.Context) {
	var input models.Author
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil membuat penulis baru",
		"data":    input,
	})
}

// UpdateAuthor - Memperbarui data penulis
func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	if err := config.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Penulis tidak ditemukan"})
		return
	}

	var input models.Author
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&author).Updates(models.Author{
		Name:   input.Name,
		Email:  input.Email,
		Bio:    input.Bio,
		Avatar: input.Avatar,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil memperbarui data penulis",
		"data":    author,
	})
}

// DeleteAuthor - Menghapus penulis
func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	if err := config.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Penulis tidak ditemukan"})
		return
	}

	config.DB.Delete(&author)
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menghapus penulis",
	})
}
