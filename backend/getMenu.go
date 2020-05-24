package main

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

type pickupMenu struct {
	gorm.Model
	ID uint
	Name string
	Price uint
	Describe string
	Genre uint
}

func gormConnect() *gorm.DB {
	DBtype := "mysql"
	user := "tkaji"
	pass := "argentina7"
	protocol := "tcp(:8080)"
	DBname := "menu"

	connect := user+":"+pass+"@"+protocol+"/"+DBname+"?parseTime=True"
	db, err := gorm.Open(DBtype, connect)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func getMenu() []pickupMenu{
	db := gormConnect()

	defer db.Close()

	var showMenu []pickupMenu
	db.Order("Genre desc").Find(&showMenu)
	return showMenu
}