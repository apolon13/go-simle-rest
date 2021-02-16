package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(c *Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(c.dsnString()), &gorm.Config{})
}
