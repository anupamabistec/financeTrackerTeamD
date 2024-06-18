package routes

import (
	"github.com/anupamabistec/financeTracker/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateExpense(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	expense := new(models.Expense)

	if err := c.BodyParser(expense); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db.Create(&expense)
	return c.JSON(expense)
}

func GetExpenses(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var expenses []models.Expense

	db.Find(&expenses)
	return c.JSON(expenses)
}

func RegisterExpenseRoutes(app *fiber.App, db *gorm.DB) {
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Post("/expenses", CreateExpense)
	app.Get("/expenses", GetExpenses)
}
