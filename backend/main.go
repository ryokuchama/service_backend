package main

import (
	"sync"
	"github.com/ant0ine/go-json-rest/rest"
    "log"
    "net/http"
)

type menu struct {
	id int
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

var store = map[string] *menu{}
var lock = sync.RWMutex{}

func getAllMenu(w rest.ResponseWriter, r *rest.Request) {
	
}