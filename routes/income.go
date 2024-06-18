package routes

import (
	"github.com/anupamabistec/financeTracker/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateIncome(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	income := new(models.Income)

	if err := c.BodyParser(income); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db.Create(&income)
	return c.JSON(income)
}

func GetIncomes(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var incomes []models.Income

	db.Find(&incomes)
	return c.JSON(incomes)
}

func RegisterIncomeRoutes(app *fiber.App, db *gorm.DB) {
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Post("/incomes", CreateIncome)
	app.Get("/incomes", GetIncomes)
}
