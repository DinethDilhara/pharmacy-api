package handlers

import (
	"regexp"
	"strconv"

	"github.com/DinethDilhara/pharmacy-api/database"
	"github.com/DinethDilhara/pharmacy-api/models"
	"github.com/gofiber/fiber/v2"
)


var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
var mobileRegex = regexp.MustCompile(`^\d{10}$`)


func GetAllInvoices(c *fiber.Ctx) error {
    var invoices []models.Invoice
    if err := database.DB.Db.Preload("Items.Item").Find(&invoices).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch invoices"})
    }
    return c.Status(fiber.StatusOK).JSON(invoices)
}


func GetInvoiceByID(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid invoice ID"})
    }

    var invoice models.Invoice
    if err := database.DB.Db.Preload("Items.Item").First(&invoice, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invoice not found"})
    }

    return c.Status(fiber.StatusOK).JSON(invoice)
}


func CreateInvoice(c *fiber.Ctx) error {
    var input models.Invoice
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    if input.ClientName == "" || !mobileRegex.MatchString(input.MobileNo) || (input.Email != "" && !emailRegex.MatchString(input.Email)) {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid client data: name required, email must be valid, and mobile number must be 10 digits"})
    }

    if err := database.DB.Db.Create(&input).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create invoice"})
    }

    return c.Status(fiber.StatusCreated).JSON(input)
}


func UpdateInvoice(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid invoice ID"})
    }

    var invoice models.Invoice
    if err := database.DB.Db.First(&invoice, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invoice not found"})
    }

    var input models.Invoice
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    if input.ClientName == "" || !mobileRegex.MatchString(input.MobileNo) || (input.Email != "" && !emailRegex.MatchString(input.Email)) {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid client data"})
    }

    invoice.ClientName = input.ClientName
    invoice.MobileNo = input.MobileNo
    invoice.Email = input.Email
    invoice.Address = input.Address
    invoice.BillingType = input.BillingType

    if err := database.DB.Db.Save(&invoice).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update invoice"})
    }

    return c.Status(fiber.StatusOK).JSON(invoice)
}


func DeleteInvoice(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid invoice ID"})
    }

    var invoice models.Invoice
    if err := database.DB.Db.First(&invoice, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invoice not found"})
    }

    if err := database.DB.Db.Delete(&invoice).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete invoice"})
    }

    return c.SendStatus(fiber.StatusNoContent)
}

func AddInvoiceItem(c *fiber.Ctx) error {
    invoiceID, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid invoice ID"})
    }

    var invoice models.Invoice
    if err := database.DB.Db.First(&invoice, invoiceID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invoice not found"})
    }

    var input models.InvoiceItem
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    if input.ItemID == 0 || input.Quantity <= 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Item ID must be provided and quantity must be > 0"})
    }

    var item models.Item
    if err := database.DB.Db.First(&item, input.ItemID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Referenced item not found"})
    }

    input.InvoiceID = uint(invoiceID)

    if err := database.DB.Db.Create(&input).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add invoice item"})
    }

    var createdItem models.InvoiceItem
    if err := database.DB.Db.Preload("Item").First(&createdItem, input.ID).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to load created invoice item"})
    }

    return c.Status(fiber.StatusCreated).JSON(createdItem)
}

func UpdateInvoiceItem(c *fiber.Ctx) error {
    invoiceID, err := strconv.Atoi(c.Params("id"))
    if err != nil || invoiceID <= 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid invoice ID"})
    }

    var input models.InvoiceItem
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    if input.ID == 0 || input.ItemID == 0 || input.Quantity <= 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invoice item ID, valid Item ID, and quantity (> 0) are required",
        })
    }

    var invoiceItem models.InvoiceItem
    if err := database.DB.Db.First(&invoiceItem, input.ID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invoice item not found"})
    }

    if invoiceItem.InvoiceID != uint(invoiceID) {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invoice item does not belong to the specified invoice",
        })
    }

    var item models.Item
    if err := database.DB.Db.First(&item, input.ItemID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Provided Item ID does not exist in the items list",
        })
    }

    invoiceItem.ItemID = input.ItemID
    invoiceItem.Quantity = input.Quantity

    if err := database.DB.Db.Save(&invoiceItem).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update invoice item",
        })
    }

    var updatedItem models.InvoiceItem
    if err := database.DB.Db.Preload("Item").First(&updatedItem, invoiceItem.ID).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to load updated invoice item details",
        })
    }

    return c.Status(fiber.StatusOK).JSON(updatedItem)
}

func DeleteInvoiceItem(c *fiber.Ctx) error {
    invoiceID, err := strconv.Atoi(c.Params("id"))
    itemID, err2 := strconv.Atoi(c.Query("item_id")) 
    if err != nil || err2 != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid invoice ID or item ID"})
    }

    var invoiceItem models.InvoiceItem
    if err := database.DB.Db.
        Where("invoice_id = ? AND id = ?", invoiceID, itemID).
        First(&invoiceItem).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invoice item not found"})
    }

    if err := database.DB.Db.Delete(&invoiceItem).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete invoice item"})
    }

    return c.SendStatus(fiber.StatusNoContent)
}
