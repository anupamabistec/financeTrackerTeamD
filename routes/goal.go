package routes

import (
	"github.com/anupamabistec/financeTracker/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateGoal(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	goal := new(models.Goal)

	if err := c.BodyParser(goal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db.Create(&goal)
	return c.JSON(goal)
}

func GetGoals(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var goals []models.Goal

	db.Find(&goals)
	return c.JSON(goals)
}

func RegisterGoalRoutes(app *fiber.App, db *gorm.DB) {
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Post("/goals", CreateGoal)
	app.Get("/goals", GetGoals)
}
