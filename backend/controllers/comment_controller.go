package controllers

import (
	"github.com/blog-cms/backend/config"
	"github.com/blog-cms/backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllComments(c *fiber.Ctx) error {
	var comments []models.Comment

	query := config.DB.Preload("Article")

	if articleID := c.Query("article_id"); articleID != "" {
		query = query.Where("article_id = ?", articleID)
	}

	result := query.Order("created_at DESC").Find(&comments)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data komentar",
		"data":    comments,
	})
}

func GetCommentByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var comment models.Comment
	result := config.DB.Preload("Article").First(&comment, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Komentar tidak ditemukan"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data komentar",
		"data":    comment,
	})
}

// CreateComment - Membuat komentar baru
func CreateComment(c *fiber.Ctx) error {
	var input models.Comment
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Pastikan artikel ada
	var article models.Article
	if err := config.DB.First(&article, input.ArticleID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Artikel tidak ditemukan"})
	}

	result := config.DB.Create(&input)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	// Reload dengan relasi
	config.DB.Preload("Article").First(&input, input.ID)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Berhasil membuat komentar baru",
		"data":    input,
	})
}

// UpdateComment - Memperbarui data komentar
func UpdateComment(c *fiber.Ctx) error {
	id := c.Params("id")
	var comment models.Comment
	if err := config.DB.First(&comment, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Komentar tidak ditemukan"})
	}

	var input models.Comment
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Model(&comment).Updates(models.Comment{
		CommenterName:  input.CommenterName,
		CommenterEmail: input.CommenterEmail,
		Content:        input.Content,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil memperbarui data komentar",
		"data":    comment,
	})
}

// DeleteComment - Menghapus komentar
func DeleteComment(c *fiber.Ctx) error {
	id := c.Params("id")
	var comment models.Comment
	if err := config.DB.First(&comment, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Komentar tidak ditemukan"})
	}

	config.DB.Delete(&comment)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menghapus komentar",
	})
}

// GetCommentsByArticleID - Mengambil komentar berdasarkan Article ID
func GetCommentsByArticleID(c *fiber.Ctx) error {
	articleID := c.Params("id")
	var comments []models.Comment
	result := config.DB.Where("article_id = ?", articleID).Order("created_at DESC").Find(&comments)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil komentar artikel",
		"data":    comments,
	})
}
