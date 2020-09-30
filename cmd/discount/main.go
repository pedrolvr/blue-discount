package main

import (
	"blue-discount/internal/app"
	"blue-discount/internal/interface/transport/grpc"
	"blue-discount/pkg/db"
	"blue-discount/tool"
	"log"

	"gorm.io/gorm"
)

func readConfig() app.Config {
	config, err := app.ReadConfig("app", "./config/")

	if err != nil {
		log.Fatalf("Read config error: %s", err)
	}

	return config
}

func connectDB(c app.DBConfig) *gorm.DB {
	dataSource := db.ConnectionString(c.Host, c.Port, c.User, c.Password, c.Name)

	conn, err := db.Connect(dataSource)

	if err != nil {
		log.Fatalf("Connect db error: %s", err)
	}

	return conn
}

func main() {
	config := readConfig()
	dbConn := connectDB(config.DB)

	sqlDB, err := dbConn.DB()

	if err != nil {
		log.Fatalf("Get sql db from gorm error: %s", err)
	}

	err = tool.MigrationUp(sqlDB)

	if err != nil {
		log.Fatalf("Migrations error: %s", err)
	}

	grpc.Start(dbConn, config)
}
