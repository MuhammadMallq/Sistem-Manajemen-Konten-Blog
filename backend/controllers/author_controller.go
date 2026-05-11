package controllers

import (
	"github.com/blog-cms/backend/config"
	"github.com/blog-cms/backend/models"
	"github.com/gofiber/fiber/v2"
)

// GetAllAuthors - Mengambil semua data penulis
func GetAllAuthors(c *fiber.Ctx) error {
	var authors []models.Author
	result := config.DB.Preload("Articles").Find(&authors)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data penulis",
		"data":    authors,
	})
}

// GetAuthorByID - Mengambil data penulis berdasarkan ID
func GetAuthorByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author
	result := config.DB.Preload("Articles").First(&author, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Penulis tidak ditemukan"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data penulis",
		"data":    author,
	})
}

// CreateAuthor - Membuat penulis baru
func CreateAuthor(c *fiber.Ctx) error {
	var input models.Author
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	result := config.DB.Create(&input)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Berhasil membuat penulis baru",
		"data":    input,
	})
}

// UpdateAuthor - Memperbarui data penulis
func UpdateAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author
	if err := config.DB.First(&author, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Penulis tidak ditemukan"})
	}

	var input models.Author
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Model(&author).Updates(models.Author{
		Name:   input.Name,
		Email:  input.Email,
		Bio:    input.Bio,
		Avatar: input.Avatar,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil memperbarui data penulis",
		"data":    author,
	})
}

// DeleteAuthor - Menghapus penulis
func DeleteAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author
	if err := config.DB.First(&author, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Penulis tidak ditemukan"})
	}

	config.DB.Delete(&author)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menghapus penulis",
	})
}
