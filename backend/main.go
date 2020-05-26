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

func postToJSON() {
	router := gin.Default()
	router.GET("/getMenu", func(c *gin.Context){
		var json jsonRequest
		if err := c.ShouldBindJSON(&json); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"str": json.FieldStr, "int": json.FieldInt, "bool": json.FieldBool})
		write()
	})
}

func writeOrder () {
	r := gin.Default()
	r.POST("/write", func(c * gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Printf("Server Side")
	r.Run()
}