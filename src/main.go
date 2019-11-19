package main

import (
	"fmt"
	"log"
	"net/http"
	"research/tripyuk/src/common/config"
	"research/tripyuk/src/infra/mysql"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

var (
	configuration config.Configuration
	dbFactory     *mysql.DBFactory
)

func setupRouter() *gin.Engine {
	db, err := dbFactory.DBConnection()
	if err != nil {
		log.Fatalf("Failed to open database connection: %s", err)
		panic(fmt.Errorf("Fatal error connecting to database: %s", err))
	}
	fmt.Println("running db")

	defer db.Close()
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func init() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
		panic(fmt.Errorf("Fatal error loading configuration: %s", err))
	}

	configuration = *cfg
	dbFactory = mysql.NewDbFactory(configuration.Database)
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
