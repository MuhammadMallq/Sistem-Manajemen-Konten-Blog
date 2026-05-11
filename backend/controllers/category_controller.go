package controllers

import (
	"github.com/blog-cms/backend/config"
	"github.com/blog-cms/backend/models"
	"github.com/gofiber/fiber/v2"
)

// GetAllCategories - Mengambil semua data kategori
func GetAllCategories(c *fiber.Ctx) error {
	var categories []models.Category
	result := config.DB.Find(&categories)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data kategori",
		"data":    categories,
	})
}

// GetCategoryByID - Mengambil data kategori berdasarkan ID
func GetCategoryByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	result := config.DB.Preload("Articles").First(&category, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Kategori tidak ditemukan"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data kategori",
		"data":    category,
	})
}

// CreateCategory - Membuat kategori baru
func CreateCategory(c *fiber.Ctx) error {
	var input models.Category
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	result := config.DB.Create(&input)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Berhasil membuat kategori baru",
		"data":    input,
	})
}

// UpdateCategory - Memperbarui data kategori
func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Kategori tidak ditemukan"})
	}

	var input models.Category
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Model(&category).Updates(models.Category{
		Name:        input.Name,
		Description: input.Description,
		Color:       input.Color,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil memperbarui data kategori",
		"data":    category,
	})
}

// DeleteCategory - Menghapus kategori
func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Kategori tidak ditemukan"})
	}

	config.DB.Delete(&category)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menghapus kategori",
	})
}
