package models

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Description string
	Amount      float64
	Category    string
	Date        string
}

type Income struct {
	gorm.Model
	Source string
	Amount float64
	Date   string
}

type Goal struct {
	gorm.Model
	Name          string
	TargetAmount  float64
	CurrentAmount float64
	DueDate       string
}
