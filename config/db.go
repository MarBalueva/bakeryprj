package config

import (
	"bakeryapp/logger"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	pgDriver "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	//dsn := "host=db user=postgres password=1 dbname=bakery port=5432 sslmode=disable search_path=bakery"
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
		logger.Log.Error("DB connection error: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Log.Error("Failed to get sql.DB from gorm.DB: ", err)
	}

	driver, err := pgDriver.WithInstance(sqlDB, &pgDriver.Config{})
	if err != nil {
		logger.Log.Error("Failed to create DB driver: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		logger.Log.Error("Migration init error: ", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Log.Error("Migration failed: ", err)
	} else if err == migrate.ErrNoChange {
		logger.Log.Info("No new migrations to apply")
	} else {
		logger.Log.Info("Migrations applied successfully")
	}

	logger.Log.Info("DB connected and migrated successfully")
	return db
}
