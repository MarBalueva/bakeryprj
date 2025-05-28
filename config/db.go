package config

import (
	"bakeryapp/logger"
	"bakeryapp/models"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable search_path=bakery",
		Cfg.Database.Host,
		Cfg.Database.User,
		Cfg.Database.Password,
		Cfg.Database.Name,
		Cfg.Database.Port)

	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		logger.Log.Errorf("DB connection attempt %d failed: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		logger.Log.Fatal("DB connection error: ", err)
	}
	//db.Exec("CREATE SCHEMA IF NOT EXISTS bakery")

	err = db.AutoMigrate(
		&models.Appuser{},
		&models.Product{},
		&models.Order{},
		&models.Employee{},
		&models.ProductInBasket{},
		&models.Client{},
		&models.Payment{},
		//&models.DictionaryItem{},
		&models.Document{},
		&models.PriceHistory{},
		&models.ProductInOrder{},
		&models.AccessGroup{},
		&models.PaymentType{},
		&models.Position{},
		&models.Status{},
		&models.SubcategoryProduct{},
		&models.UserAccess{},
	)
	if err != nil {
		logger.Log.Fatal("Auto migration failed: ", err)
	}

	logger.Log.Info("DB connected and schema auto-migrated")
	return db
}
