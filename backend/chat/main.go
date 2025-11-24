package main

import (
	"fmt"
	"net/http"

	"chat/db"
	"chat/pets"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // postgres driver
)

func main() {
	fmt.Println("Chat App")

	db := db.ConnectDB()
	defer db.Close()

	router := gin.Default() // Creates a Gin router
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "home"})
	})
	pets.RegisterPetsRoutes(router, db)

	router.Run("localhost:6969")
}
