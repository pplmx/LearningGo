package database

import (
	"fmt"
	"github.com/pplmx/LearningGo/fiber_boot/app/models"
	"github.com/pplmx/LearningGo/fiber_boot/pkg/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		env.GetEnv("DB_HOST", "127.0.0.1"),
		env.GetEnv("DB_USER", ""),
		env.GetEnv("DB_PASSWORD", ""),
		env.GetEnv("DB_NAME", ""),
		env.GetEnv("DB_PORT", "5432"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Printf("Error migrating database: %s\n", err.Error())
		return
	}
}
