package infra

import (
	"log"

	"github.com/yakan15/clean-transaction/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=user password=pass dbname=user port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func AutoMigrate(d *gorm.DB) {
	if err := d.AutoMigrate(
		&model.Article{},
	); err != nil {
		log.Fatal(err)
	}
}
