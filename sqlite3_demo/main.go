package main

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	dsn := "file:./x.db?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)"
	// dsn := "file::memory:?cache=shared&_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect sqlite3 database: %v\n", err)
		return
	}

	// create table
	db.Exec("create table if not exists t1 (id int, name varchar(20))")
	db.Exec("delete from t1")

	// write data in concurrent by goroutine
	for i := 0; i < 1000; i++ {
		go func(i int) {
			db.Exec("insert into t1 (id, name) values (?, ?)", i, fmt.Sprintf("name%d", i))
		}(i)
	}

	// update data in concurrent by goroutine
	for i := 0; i < 1000; i++ {
		go func(i int) {
			db.Exec("update t1 set name = ? where id = ?", fmt.Sprintf("name%d", i), i)
		}(i)
	}

}
