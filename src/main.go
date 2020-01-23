package main

import (
	"fmt"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tripyuk/src/common/config"
	"tripyuk/src/common/middleware"
	"tripyuk/src/infra/mysql"
)

var db = make(map[string]string)

var (
	configuration config.Configuration
	dbFactory     *mysql.DBFactory
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

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


	defaultMiddleware, err := middleware.DefaultMW()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}


	r.POST("/login", defaultMiddleware.LoginHandler)

	r.NoRoute(defaultMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", defaultMiddleware.RefreshHandler)
	auth.Use(defaultMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", middleware.HelloHandler)
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
