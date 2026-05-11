package controllers

import (
	"strings"
	"time"

	"github.com/blog-cms/backend/config"
	"github.com/blog-cms/backend/models"
	"github.com/gofiber/fiber/v2"
)

// GetAllArticles - Mengambil semua data artikel dengan relasi
func GetAllArticles(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data artikel",
		"data":    articles,
		"total":   len(articles),
	})
}

// GetArticleByID - Mengambil data artikel berdasarkan ID
func GetArticleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var article models.Article
	result := config.DB.Preload("Author").Preload("Category").Preload("Comments").First(&article, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Artikel tidak ditemukan"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data artikel",
		"data":    article,
	})
}

// CreateArticle - Membuat artikel baru
func CreateArticle(c *fiber.Ctx) error {
	var input models.Article
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	// Reload dengan relasi
	config.DB.Preload("Author").Preload("Category").First(&input, input.ID)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Berhasil membuat artikel baru",
		"data":    input,
	})
}

// UpdateArticle - Memperbarui data artikel
func UpdateArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article models.Article
	if err := config.DB.First(&article, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Artikel tidak ditemukan"})
	}

	var input models.Article
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil memperbarui data artikel",
		"data":    article,
	})
}

// DeleteArticle - Menghapus artikel
func DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article models.Article
	if err := config.DB.First(&article, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Artikel tidak ditemukan"})
	}

	// Hapus komentar terkait terlebih dahulu
	config.DB.Where("article_id = ?", id).Delete(&models.Comment{})
	config.DB.Delete(&article)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menghapus artikel",
	})
}

// GetDashboardStats - Mengambil statistik dashboard
func GetDashboardStats(c *fiber.Ctx) error {
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil statistik dashboard",
		"data": fiber.Map{
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
