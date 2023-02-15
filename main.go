package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"go-web-demo/database"
	"go-web-demo/redisgo"
	"net/http"
)

type album struct {
	ID     string  `json:id`
	Title  string  `json:title`
	Artist string  `json:artist`
	Price  float64 `json:price`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jerry", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func selectUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := database.SelectById(id)
	if err != nil {
		fmt.Println("user select err,", err)
	} else {
		c.IndentedJSON(http.StatusOK, user)
	}
}

func cacheKeyValue(c *gin.Context) {
	data, _ := c.GetRawData()
	var m map[string]interface{}
	_ = json.Unmarshal(data, &m)

	key, value := m["key"], m["value"]

	conn := redisgo.GetConn()
	_, err1 := conn.Do("set", key, value)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, "key value cache failed!")
		return
	}
	values, _ := redis.String(conn.Do("get", key))
	fmt.Println(values)
	c.IndentedJSON(http.StatusOK, "success!")
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/addAlbums", postAlbums)

	// DB使用
	router.GET("/user/:id", selectUserById)

	// Redis使用
	router.POST("/cache", cacheKeyValue)

	router.Run("localhost:8080")
}
