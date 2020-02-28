package router

import (
	"log"
	"net/http"
	"tripyuk_api_go/src/common/middleware"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func RouteInit(db *gorm.DB, port string) {
	r := setupRouter()

	defaultMiddleware, err := middleware.DefaultMW(db)
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

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
