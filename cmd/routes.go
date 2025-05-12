package main

import (
	"github.com/DinethDilhara/pharmacy-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

    app.Get("/items", handlers.ListItems)
    app.Get("/items/:id", handlers.GetItemByID)
    app.Post("/items", handlers.CreateItem)
    app.Put("/items/:id", handlers.UpdateItem)
    app.Delete("/items/:id", handlers.DeleteItem)

    app.Get("/invoices/", handlers.GetAllInvoices)               
    app.Get("/invoices/:id", handlers.GetInvoiceByID)            
    app.Post("/invoices/", handlers.CreateInvoice)               
    app.Put("/invoices/:id", handlers.UpdateInvoice)             
    app.Delete("/invoices/:id", handlers.DeleteInvoice)          

    app.Post("/invoices/:id/items", handlers.AddInvoiceItem)     
    app.Put("/invoices/:id/items", handlers.UpdateInvoiceItem)    
    app.Delete("/invoices/:id/items", handlers.DeleteInvoiceItem)
}


