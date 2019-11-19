package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"research/tripyuk/src/common/config"
)

// DBFactory struct
type DBFactory struct {
	config config.DatabaseConfiguration
}

// NewDbFactory instantiate new DB Factory object
func NewDbFactory(cfg config.DatabaseConfiguration) *DBFactory {
	return &DBFactory{config: cfg}
}

// DBConnection get open database connection
func (f *DBFactory) DBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(f.config.DbType, f.config.ConnectionUri)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
		return nil, err
	}
	return db, nil
}
