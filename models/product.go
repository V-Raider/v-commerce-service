package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Category    string
	Description string
	// Image
}
