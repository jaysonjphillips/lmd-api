package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Conn *gorm.DB
)

func init() {
	var err error
	Conn, err = gorm.Open(sqlite.Open("dashboard.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}
