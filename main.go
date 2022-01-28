package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db map[string]string

func main() {
	db = make(map[string]string)

	router := gin.Default()
	router.GET("/:key", getData)
	router.GET("/", getAll)
	router.PUT("/:key/:value", putData)
	router.DELETE("/:key", deleteData)
	router.Run("localhost:8080")
}

func getAll(c *gin.Context) {
	c.JSON(http.StatusOK, db)
}

func getData(c *gin.Context) {
	data, ok := db[c.Param("key")]
	if ok {
		c.JSON(http.StatusOK, gin.H{c.Param("key"): data})
	}
}

func putData(c *gin.Context) {
	db[c.Param("key")] = c.Param("value")
}

func deleteData(c *gin.Context) {
	data, ok := db[c.Param("key")]
	if ok {
		c.JSON(http.StatusOK, gin.H{c.Param("key"): data})
		delete(db, c.Param("key"))
	}
}
