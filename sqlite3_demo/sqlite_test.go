package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	UID string `gorm:"index:idx_user_uid,unique"`
}

func BenchmarkCreateUsers(b *testing.B) {
	db, err := initDB()
	require.Nil(b, err)

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
	db, err := initDB()
	require.Nil(b, err)
	db.Create(&User{UID: "uid_1"})

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				uuidStr := uuid.NewString()
				db.Model(&User{}).Where("id = ?", 1).Update("uid", "uid_"+uuidStr)
			}
		},
	)
}

func initDB() (*gorm.DB, error) {
	// dsn := "file:./x.db?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)"
	dsn := "file::memory:?cache=shared&_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)"

	lg := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{LogLevel: logger.Error})
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: lg,
	})
	if err != nil {
		fmt.Printf("failed to connect sqlite3 database: %v\n", err)
		return nil, err
	}

	err = db.Migrator().DropTable(&User{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
