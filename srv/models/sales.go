package models

import (
	"gorm.io/gorm"
)

type Sales struct {
	gorm.Model
	UserID      uint // Fk to User
	User        User `gorm:"foreignKey:UserID"` // Relation with User
	Description string
	Total       float64
}

func CreateSale(db *gorm.DB, sale *Sales) error {
	result := db.Create(sale)
	return result.Error
}

func GetAllSales(db *gorm.DB) ([]Sales, error) {
	var sales []Sales
	result := db.Preload("User").Find(&sales) // Preloads the associated user data.
	return sales, result.Error
}
