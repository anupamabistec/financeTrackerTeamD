package main

import (
	"log"

	"github.com/anupamabistec/financeTracker/models" // Correct import path
	"github.com/anupamabistec/financeTracker/routes" // Correct import path
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize the Fiber app
	app := fiber.New()

	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("finance.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Auto-migrate models
	db.AutoMigrate(&models.Expense{}, &models.Income{}, &models.Goal{})

	// Register routes
	routes.RegisterExpenseRoutes(app, db)
	routes.RegisterIncomeRoutes(app, db)
	routes.RegisterGoalRoutes(app, db)
	routes.RegisterReportRoutes(app, db)

	// Start the Fiber app
	log.Fatal(app.Listen(":3000"))
}
