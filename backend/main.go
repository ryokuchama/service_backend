package main

import (
	"sync"
	"github.com/ant0ine/go-json-rest/rest"
    "log"
	"net/http"

)

// メニュー構造体
type menu struct {
	Id int
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

// メニュー取得
func getAllMenu(w rest.ResponseWriter, r *rest.Request) {
	
	lock.RLock()
	menus := make([]menu, len(store))
	i := 0
	for _, m := range store {
		menus[i] = *m
		i++
	}
	lock.RUnlock()
	w.WriteJson(&menus)
}

// メニュー追加
func addMenu (w rest.ResponseWriter, r *rest.Request) {
	menus := menu{}
	err := r.DecodeJsonPayload(&menus)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lock.Lock()
	store[menus.Id] = &menus
	lock.Unlock()
	w.WriteJson(&menus)
}

// メニュー削除
func deleteMenu (w rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("Id")
	lock.Lock()
	delete(store, code)
	lock.Unlock()
	w.WriteHeader(http.StatusOK)
}

var storeOrder = map[int] *order{}
var lockOrder = sync.RWMutex{}

// 注文の構造体
type order struct {
	Id int
	Name string
	Amount int
	Price int
	Date string
}
// オーダー登録
func write(w rest.ResponseWriter, r *rest.Request){
	orders := order{}
	err := r.DecodeJsonPayload(&orders)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	lock.Lock()
	storeOrder[orders.Id] = &orders
	lockOrder.Unlock()
	w.WriteJson(&orders)
}