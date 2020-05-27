package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type menu struct {
	id int
	name string
	price int
	text string
}

func main () {
	router := gin.Default()
	router.GET("/getmenu", func(c*gin.Context) {
		pickupmenu := getMenu()
		c.JSON(200, gin.H{
			"menu" : pickupmenu,
		})
	})
	router.POST("/insert", func(c*gin.Context) {
		var form menu

		if err := c.Bind(&form); err != nil {
			pickupmenu := getMenu()
			c.JSON(200, gin.H{
				"menu" : pickupmenu,
			})
		}
	})
	router.PUT("/update", func(c*gin.Context) {

	})
	router.DELETE("/delete")

	router.POST("/order")

	router.Run()
}