package main

import (
	"fmt"
	"github.com/marcosstupnicki/go-users/internal/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

const (
	ExitCodeFailToConnectLocalDB = iota
	ExitCodeFailToMigrateModel
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	user := "root"
	password := "root"
	host := "127.0.0.1"
	port := "3306"
	dbname := "users"


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Sprintln("Unable to connect local db")
		os.Exit(ExitCodeFailToConnectLocalDB)
	}

	err = db.AutoMigrate(&users.User{})
	if err != nil {
		os.Exit(ExitCodeFailToMigrateModel)
	}
}
