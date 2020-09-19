package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionString(host string, port int, username string, password string, dbName string) string {
    mask := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
    return fmt.Sprintf(mask, host, port, username, password, dbName)
}

func Connect(dsn string) (db *gorm.DB, err error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
