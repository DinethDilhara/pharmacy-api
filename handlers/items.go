package handlers

import (
	"strconv"

	"github.com/DinethDilhara/pharmacy-api/database"
	"github.com/DinethDilhara/pharmacy-api/models"
	"github.com/gofiber/fiber/v2"
)

func ListItems(c *fiber.Ctx) error {
    var items []models.Item
    if err := database.DB.Db.Find(&items).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch items"})
    }
    return c.Status(fiber.StatusOK).JSON(items)
}


func GetItemByID(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid item ID"})
    }

    var item models.Item
    if err := database.DB.Db.First(&item, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
    }

    return c.Status(fiber.StatusOK).JSON(item)
}


func CreateItem(c *fiber.Ctx) error {
    var input models.Item
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    if input.Name == "" || input.Category == "" || input.UnitPrice <= 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Name and category must be non-empty, and unit_price must be greater than 0",
        })
    }

    if err := database.DB.Db.Create(&input).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create item"})
    }

    return c.Status(fiber.StatusCreated).JSON(input)
}


func UpdateItem(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid item ID"})
    }

    var item models.Item
    if err := database.DB.Db.First(&item, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
    }

    var input models.Item
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    if input.Name == "" || input.Category == "" || input.UnitPrice <= 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Name and category must be non-empty, and unit_price must be greater than 0",
        })
    }

    item.Name = input.Name
    item.Category = input.Category
    item.UnitPrice = input.UnitPrice

    if err := database.DB.Db.Save(&item).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update item"})
    }

    return c.Status(fiber.StatusOK).JSON(item)
}

func DeleteItem(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid item ID"})
    }

    var item models.Item
    if err := database.DB.Db.First(&item, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
    }

    if err := database.DB.Db.Delete(&item).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete item"})
    }

    return c.SendStatus(fiber.StatusNoContent)
}
