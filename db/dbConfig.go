package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	username string
	password string
	dbName   string
	dbHost   string
	dbPort   string
	charset  string
}

func (c *Config) dsnString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		c.username,
		c.password,
		c.dbHost,
		c.dbPort,
		c.dbName,
		c.charset,
	)
}

func NewConfig() (*Config, error) {
	if e := godotenv.Load(); e != nil {
		return nil, e
	}
	return &Config{
		os.Getenv("db_user"),
		os.Getenv("db_pass"),
		os.Getenv("db_name"),
		os.Getenv("db_host"),
		os.Getenv("db_port"),
		os.Getenv("charset"),
	}, nil
}
