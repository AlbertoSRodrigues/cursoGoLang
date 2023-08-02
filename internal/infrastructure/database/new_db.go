package database

import (
	"projeto/internal/domain/campaign"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {

	dsn := "host=localhost user=admin password=root dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database.")
	}
	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})
	return db
}
