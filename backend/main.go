package main

import (
	"strconv"
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
		n := c.Param("id")
		id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
		}
		updatemenu := c.PostForm("getmenu")

	})
	router.DELETE("/delete", func(c*gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
	})

	router.POST("/order")

	router.Run()
}