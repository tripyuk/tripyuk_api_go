package main

import (
	"fmt"
	"log"
	"tripyuk_api_go/src/common/config"
	"tripyuk_api_go/src/infra/mysql"
	"tripyuk_api_go/src/router"
)

var db = make(map[string]string)

var (
	configuration config.Configuration
	dbFactory     *mysql.DBFactory
	port          string
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func init() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
		panic(fmt.Errorf("Fatal error loading configuration: %s", err))
	}

	configuration = *cfg
	dbFactory = mysql.NewDbFactory(configuration.Database)
	port = configuration.Server.Port
}

func main() {

	db, err := dbFactory.DBConnection()
	if err != nil {
		log.Fatalf("Failed to open database connection: %s", err)
		panic(fmt.Errorf("Fatal error connecting to database: %s", err))
	}
	fmt.Println("running db")

	defer db.Close()

	router.RouteInit(db, port)
	// Listen and Server in 0.0.0.0:8080
}
