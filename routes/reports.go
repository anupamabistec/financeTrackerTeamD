package routes

import (
	"github.com/anupamabistec/financeTracker/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GenerateExpenseReport(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var expenses []models.Expense

	db.Find(&expenses)
	//reports.GenerateExpenseReport(expenses)
	return c.SendString("Expense report generated: expense_report.html")
}

func RegisterReportRoutes(app *fiber.App, db *gorm.DB) {
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Get("/reports/expenses", GenerateExpenseReport)
}
