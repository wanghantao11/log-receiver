package database

import (
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/wanghantao11/log-receiver/config"
)

// DBProvider is gorm.DB.
type DBProvider struct {
	gormDB *gorm.DB
}

// NewDBProvider builds a new DBProvider.
func NewDBProvider(database *gorm.DB) *DBProvider {
	dbProvider := &DBProvider{
		gormDB: database,
	}
	return dbProvider
}

// DB get a gorm DB for provider.
func (dbProvider *DBProvider) DB() *gorm.DB {
	return dbProvider.gormDB
}

// ConnectionString builds args for db connection.
func ConnectionString() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s password=%s",
		config.Get(config.DBIP),
		config.Get(config.DBUSERNAME),
		config.Get(config.DBNAME),
		config.Get(config.DBPASSWORD),
	)
}
