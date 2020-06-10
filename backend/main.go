package main

import (
	"sync"
	"github.com/ant0ine/go-json-rest/rest"
    "log"
	"net/http"
	"fmt"
	"github.com/jinzhu/gorm"
)

// メニュー構造体
type menu struct {
	gorm.Model
	ID int
	name string
	price int
	text string
}

func main () {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/getAllMenu", getAllMenu),
		rest.Post("/writeOrder", write),
		rest.Post("/addMenu", addMenu),
		rest.Delete("/deleteMenu", deleteMenu),
	)
	if err != nil {
		log.Fatal(router)
	}
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

var store = map[int] *menu{}
var lock = sync.RWMutex{}

// 全メニュー取得
func getAllMenu(w rest.ResponseWriter, r *rest.Request) {
	
	db := gormConnect()
	defer db.Close()

	menus := menu{}
	db.Find(&menus)
	fmt.Println(menus)

	w.WriteHeader(http.StatusOK)
	w.WriteJson(&menus)	
}

// メニュー追加
func addMenu (w rest.ResponseWriter, r *rest.Request) {

	db := gormConnect()
	defer db.Close()

	menus := menu{}
	
	db.NewRecord(&menus)
	db.Create(&menus)

	w.WriteJson(&menus)
}

// メニュー削除
func deleteMenu (w rest.ResponseWriter, r *rest.Request) {
	db := gormConnect()
	defer db.Close()

	menus := menu{}
	db.Delete(&menus)
}

// 注文の構造体
type order struct {
	gorm.Model
	ID int
	Name string
	Amount int
	Price int
	Date string
}

// オーダー登録
func write(w rest.ResponseWriter, r *rest.Request){

	db := gormConnect()
	defer db.Close()

	orders := order{}
	db.NewRecord(&orders)
	db.Create(&orders)
	
	w.WriteJson(&orders)
}