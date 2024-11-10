package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/FARRAS-DARKUNO/library-management/category-service/config"
	"github.com/FARRAS-DARKUNO/library-management/category-service/models"
)

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing category")
	}

	// Simpan kategori ke database
	if err := config.DB.Create(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create category")
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve categories")
	}

	return c.Status(fiber.StatusOK).JSON(categories)
}

func GetCategoryByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Category not found")
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Category not found")
	}

	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing category")
	}

	if err := config.DB.Save(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update category")
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Category not found")
	}

	if err := config.DB.Delete(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete category")
	}

	return c.Status(fiber.StatusOK).SendString("Category deleted")
}
