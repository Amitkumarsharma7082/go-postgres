package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ! 10.
type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

// creating a connection to a PostgreSQL database using the Gorm ORM library.
func NewConnection(config *Config) (*gorm.DB, error) {
	dsn :=
		fmt.Sprintf(
			"host=%s port=%s password=%s user=%s dbname=%s sslode=%s",
			config.Host, config.Port, config.Password, config.User, config.DBName, config.SSLMode,
		)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
