package mysql

import (
	"log"
	"tripyuk_api_go/src/common/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
