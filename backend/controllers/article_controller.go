package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/blog-cms/backend/config"
	"github.com/blog-cms/backend/models"
	"github.com/gin-gonic/gin"
)

// GetAllArticles - Mengambil semua data artikel dengan relasi
func GetAllArticles(c *gin.Context) {
	var articles []models.Article

	query := config.DB.Preload("Author").Preload("Category").Preload("Comments")

	// Filter berdasarkan status
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// Filter berdasarkan kategori
	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	// Filter berdasarkan penulis
	if authorID := c.Query("author_id"); authorID != "" {
		query = query.Where("author_id = ?", authorID)
	}

	// Pencarian berdasarkan judul
	if search := c.Query("search"); search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}

	result := query.Order("created_at DESC").Find(&articles)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data artikel",
		"data":    articles,
		"total":   len(articles),
	})
}

// GetArticleByID - Mengambil data artikel berdasarkan ID
func GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	result := config.DB.Preload("Author").Preload("Category").Preload("Comments").First(&article, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data artikel",
		"data":    article,
	})
}

// CreateArticle - Membuat artikel baru
func CreateArticle(c *gin.Context) {
	var input models.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate slug dari judul
	input.Slug = generateSlug(input.Title)

	// Jika status published, set published_at
	if input.Status == "published" {
		now := time.Now()
		input.PublishedAt = &now
	}

	// Generate excerpt dari content jika kosong
	if input.Excerpt == "" && len(input.Content) > 150 {
		input.Excerpt = input.Content[:150] + "..."
	} else if input.Excerpt == "" {
		input.Excerpt = input.Content
	}

	result := config.DB.Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Reload dengan relasi
	config.DB.Preload("Author").Preload("Category").First(&input, input.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil membuat artikel baru",
		"data":    input,
	})
}

// UpdateArticle - Memperbarui data artikel
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	var input models.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update slug jika title berubah
	if input.Title != "" {
		input.Slug = generateSlug(input.Title)
	}

	// Jika status berubah ke published, set published_at
	if input.Status == "published" && article.PublishedAt == nil {
		now := time.Now()
		input.PublishedAt = &now
	}

	config.DB.Model(&article).Updates(map[string]interface{}{
		"title":        input.Title,
		"slug":         input.Slug,
		"content":      input.Content,
		"excerpt":      input.Excerpt,
		"cover_image":  input.CoverImage,
		"status":       input.Status,
		"author_id":    input.AuthorID,
		"category_id":  input.CategoryID,
		"published_at": input.PublishedAt,
	})

	// Reload dengan relasi
	config.DB.Preload("Author").Preload("Category").Preload("Comments").First(&article, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil memperbarui data artikel",
		"data":    article,
	})
}

// DeleteArticle - Menghapus artikel
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	// Hapus komentar terkait terlebih dahulu
	config.DB.Where("article_id = ?", id).Delete(&models.Comment{})
	config.DB.Delete(&article)

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menghapus artikel",
	})
}

// GetDashboardStats - Mengambil statistik dashboard
func GetDashboardStats(c *gin.Context) {
	var totalArticles, totalAuthors, totalCategories, totalComments int64
	var publishedArticles, draftArticles int64

	config.DB.Model(&models.Article{}).Count(&totalArticles)
	config.DB.Model(&models.Author{}).Count(&totalAuthors)
	config.DB.Model(&models.Category{}).Count(&totalCategories)
	config.DB.Model(&models.Comment{}).Count(&totalComments)
	config.DB.Model(&models.Article{}).Where("status = ?", "published").Count(&publishedArticles)
	config.DB.Model(&models.Article{}).Where("status = ?", "draft").Count(&draftArticles)

	// Artikel terbaru
	var recentArticles []models.Article
	config.DB.Preload("Author").Preload("Category").
		Order("created_at DESC").Limit(5).Find(&recentArticles)

	// Komentar terbaru
	var recentComments []models.Comment
	config.DB.Preload("Article").
		Order("created_at DESC").Limit(5).Find(&recentComments)

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil statistik dashboard",
		"data": gin.H{
			"total_articles":     totalArticles,
			"total_authors":      totalAuthors,
			"total_categories":   totalCategories,
			"total_comments":     totalComments,
			"published_articles": publishedArticles,
			"draft_articles":     draftArticles,
			"recent_articles":    recentArticles,
			"recent_comments":    recentComments,
		},
	})
}

// Helper: Generate slug dari title
func generateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "'", "")
	slug = strings.ReplaceAll(slug, "\"", "")
	slug = strings.ReplaceAll(slug, ".", "")
	slug = strings.ReplaceAll(slug, ",", "")
	slug = strings.ReplaceAll(slug, "!", "")
	slug = strings.ReplaceAll(slug, "?", "")
	return slug
}
