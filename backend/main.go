package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()
	router.LoadHTMLFiles("templates/interface.html")

	router.GET("/", index)
	router.POST("/insert", postData)
	router.GET("/get", getData)

	router.Run("localhost:8080")
}