package controllers

import (
	"github.com/FARRAS-DARKUNO/library-management/author-service/config"
	"github.com/FARRAS-DARKUNO/library-management/author-service/models"
	"github.com/gofiber/fiber/v2"
)

// CreateAuthor membuat data penulis baru
func CreateAuthor(c *fiber.Ctx) error {
	var author models.Author
	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request")
	}

	if err := config.DB.Create(&author).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create author")
	}

	return c.Status(fiber.StatusOK).JSON(author)
}

// GetAuthors mengambil semua penulis
func GetAuthors(c *fiber.Ctx) error {
	var authors []models.Author
	if err := config.DB.Find(&authors).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch authors")
	}

	return c.Status(fiber.StatusOK).JSON(authors)
}

// GetAuthorByID mengambil penulis berdasarkan ID
func GetAuthorByID(c *fiber.Ctx) error {
	authorID := c.Params("id")
	var author models.Author
	if err := config.DB.First(&author, authorID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Author not found")
	}

	return c.Status(fiber.StatusOK).JSON(author)
}

// UpdateAuthor memperbarui data penulis berdasarkan ID
func UpdateAuthor(c *fiber.Ctx) error {
	authorID := c.Params("id")
	var author models.Author
	if err := config.DB.First(&author, authorID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Author not found")
	}

	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request")
	}

	if err := config.DB.Save(&author).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update author")
	}

	return c.Status(fiber.StatusOK).JSON(author)
}

// DeleteAuthor menghapus data penulis berdasarkan ID
func DeleteAuthor(c *fiber.Ctx) error {
	authorID := c.Params("id")
	if err := config.DB.Delete(&models.Author{}, authorID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete author")
	}

	return c.Status(fiber.StatusOK).SendString("Author deleted successfully")
}
