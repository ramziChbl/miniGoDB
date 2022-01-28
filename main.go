package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var db map[string]string
var m sync.RWMutex

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
	m.RLock()
	c.JSON(http.StatusOK, db)
	m.RUnlock()
}

func getData(c *gin.Context) {
	m.RLock()
	data, ok := db[c.Param("key")]
	if ok {
		c.JSON(http.StatusOK, gin.H{c.Param("key"): data})
	}
	m.RUnlock()
}

func putData(c *gin.Context) {
	m.Lock()
	db[c.Param("key")] = c.Param("value")
	m.Unlock()
	m.RLock()
	data, ok := db[c.Param("key")]
	if ok {
		c.JSON(http.StatusOK, gin.H{c.Param("key"): data})
	}
	m.RUnlock()
}

func deleteData(c *gin.Context) {
	m.RLock()
	data, ok := db[c.Param("key")]
	m.RUnlock()
	if ok {
		m.Lock()
		delete(db, c.Param("key"))
		m.Unlock()
		c.JSON(http.StatusOK, gin.H{c.Param("key"): data})
	}
}
