package pets

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func RegisterPetsRoutes(r *gin.Engine, database *sql.DB) {
	db = database
	pets := r.Group("/pets")

	pets.GET("/", get_pets)
}

// GET /pets - returns list of all pet names
func get_pets(c *gin.Context) {
	fmt.Println("called get_pets")
	rows, err := db.Query("SELECT name FROM pets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var pets []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		pets = append(pets, name)
	}

	c.JSON(http.StatusOK, gin.H{"pets": pets})
}
