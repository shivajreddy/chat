package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres driver
)

var db *sql.DB

func connectDB() *sql.DB {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DB_SSLMODE")

	// Build DSN
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot ping DB:", err)
	}

	log.Println("Successfully connected")
	return db
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "home"})
}

// GET /pets - returns list of all pet names
func get_pets(c *gin.Context) {
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

func main() {
	fmt.Println("Chat App")

	db = connectDB()
	defer db.Close()

	router := gin.Default() // Creates a Gin router
	router.GET("/", home)
	router.GET("/pets", get_pets)
	router.Run("localhost:6969")
}
