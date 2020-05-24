package main

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

type writeDB struct {
	ID uint
	Name string
	Price uint
	Count uint
	Date string
}

func gormConnection() *gorm.DB {
	DBtype := "mysql"
	user := "tkaji"
	pass := "argentina7"
	protocol := "tcp(:8080)"
	DBname := "order"

	connect := user+":"+pass+"@"+protocol+"/"+DBname+"?parseTime=True"
	db, err := gorm.Open(DBtype, connect)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func write() []writeDB{
	db := gormConnection()
	defer db.Close()

	var writeData []writeDB
	db.Order("Genre desc").Find(&writeData)
	return writeData
}