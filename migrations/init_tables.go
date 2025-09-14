package migrations

import (
	"golang-toko/models"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Product{},
		&models.ProductCategory{},
		&models.Transaction{},
		&models.TransactionItem{},
		&models.User{},
	)
}
