package controllers

import (
	"net/http"

	"github.com/blog-cms/backend/config"
	"github.com/blog-cms/backend/models"
	"github.com/gin-gonic/gin"
)

// GetAllCategories - Mengambil semua data kategori
func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	result := config.DB.Find(&categories)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data kategori",
		"data":    categories,
	})
}

// GetCategoryByID - Mengambil data kategori berdasarkan ID
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	result := config.DB.Preload("Articles").First(&category, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data kategori",
		"data":    category,
	})
}

// CreateCategory - Membuat kategori baru
func CreateCategory(c *gin.Context) {
	var input models.Category
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
		"message": "Berhasil membuat kategori baru",
		"data":    input,
	})
}

// UpdateCategory - Memperbarui data kategori
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}

	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&category).Updates(models.Category{
		Name:        input.Name,
		Description: input.Description,
		Color:       input.Color,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil memperbarui data kategori",
		"data":    category,
	})
}

// DeleteCategory - Menghapus kategori
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}

	config.DB.Delete(&category)
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menghapus kategori",
	})
}
