package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)
func getRequest(){
	r := gin.Default()
	r.GET("/ping", func(c * gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Printf("Server Side")
}