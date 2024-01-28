package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Shonobi struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Age int64 `json:"age"`
	Vilagge string `json:"vilagge"`
	Clan string `json:"clan"`
	MainJutsu string `json:"main_jutsu"`
}

var Shinobis = []Shonobi{
	{ID: 1, Name: "Naruto", Age: 17, Vilagge: "Konoha", Clan: "Uzumaki", MainJutsu: "Rasengan"},
	{ID: 2, Name: "Sasuke", Age: 17, Vilagge: "Konoha", Clan: "Uchiha", MainJutsu: "Chidori"},
	{ID: 3, Name: "Sakura", Age: 17, Vilagge: "Konoha", Clan: "Haruno", MainJutsu: "Byakugou"},
	{ID: 4, Name: "Kakashi", Age: 30, Vilagge: "Konoha", Clan: "Hatake", MainJutsu: "Raikiri"},
}

func getShinobis(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Shinobis)
}

func postShinobis(c *gin.Context) {
	var newShinobi Shonobi

	if err := c.BindJSON(&newShinobi); err != nil {
		return
	}

	Shinobis = append(Shinobis, newShinobi)
	c.IndentedJSON(http.StatusCreated, newShinobi)
}

func getShinobiByID(c *gin.Context) {
	id := c.Param("id")

	for _, shonobi := range Shinobis {
		if strconv.Itoa(int(shonobi.ID)) == id {
			c.IndentedJSON(http.StatusOK, shonobi)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "shinobi not found"})
}

func main() {
	router := gin.Default()
	router.GET("/shinobis", getShinobis)
	router.POST("/shinobis", postShinobis)
	router.GET("/shinobis/:id", getShinobiByID)
	router.Run("localhost:8080")
}
