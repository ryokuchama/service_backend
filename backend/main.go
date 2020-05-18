package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type menu struct {
	id int
	name string
	price int
	text string
}


func getRequest () {
	r := gin.Default()
	r.GET("/ping", func(c * gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Printf("Server Side")
	r.Run()
}

func getMenu () {

}

func writeReserve () {

}