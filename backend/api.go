package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "interface.html", gin.H{})
}

func postData(c *gin.Context) {
	var newData Sensor

	if err := c.BindJSON(&newData); err != nil {
		return
	}

	newData.Time = time.Now()

	status, err := InsertData(newData)

	if !status && err != nil {
		panic(err.Error())
	}
	
	c.IndentedJSON(http.StatusCreated, newData)
}

func getData(c *gin.Context) {
	data, err := Query()
	if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, data)
}