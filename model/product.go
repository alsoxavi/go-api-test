package model

import (
	"gorm.io/gorm"
)

// Product model
type Product struct {
	gorm.Model  `swaggerignore:"true"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
}
