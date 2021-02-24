package main

import (
	"app/domains"
	"app/infrastructures"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
	dsn = "root:password@tcp(127.0.0.1:3306)/go_sample?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	dbInit()
	infrastructures.Init()
}

func dbInit() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Migrator().CreateTable(domains.Todo{})
}
