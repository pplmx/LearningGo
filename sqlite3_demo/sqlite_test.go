package main

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
)

type User struct {
	gorm.Model
	UID string `gorm:"index:idx_user_uid,unique"`
}

func BenchmarkCreateUsers(b *testing.B) {
	// dsn := "file:./x.db?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)"
	dsn := "file::memory:?cache=shared&_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)"

	lg := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{LogLevel: logger.Error})
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: lg,
	})
	if err != nil {
		fmt.Printf("failed to connect sqlite3 database: %v\n", err)
		return
	}

	_ = db.AutoMigrate(&User{})

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				uuidStr := uuid.NewString()
				db.Create(&User{UID: "uid_" + uuidStr})
			}
		},
	)
}

func BenchmarkUpdateUsers(b *testing.B) {
	// dsn := "file:./x.db?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)"
	dsn := "file::memory:?cache=shared&_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)"

	lg := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{LogLevel: logger.Error})
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: lg,
	})
	if err != nil {
		fmt.Printf("failed to connect sqlite3 database: %v\n", err)
		return
	}

	_ = db.AutoMigrate(&User{})

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				uuidStr := uuid.NewString()
				db.Model(&User{}).Where("id = ?", 1).Update("uid", "uid_"+uuidStr)
			}
		},
	)
}
