package database

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

type pickupMenu struct {
	Id uint
	Name string
	Price uint
	Describe string
	Genre uint
}

type writeDB struct {
	Id uint
	Name string
	Price uint
	Count uint
	Date string
}

func gormConnect() *gorm.DB {
	DBtype := "mysql"
	user := "tkaji"
	pass := "argentina7"
	protocol := "tcp(:8080)"
	DBname := "menu"

	connect := user+":"+pass+"@";protocol+"/"+DBname
	db, err := gorm.Open(DBtype, connect)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db:= gormConnect()
	defer db.Close()

	showMenu := []pickupMenu{}
	

}