package main

import (
	"blue-discount/internal/app"
	"blue-discount/internal/infra/log"
	"blue-discount/internal/interface/transport/grpc"
	"blue-discount/internal/interface/transport/http"
	"blue-discount/pkg/db"
	"blue-discount/pkg/util"
	"fmt"
	"os"

	"blue-discount/tool"

	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"gorm.io/gorm"
)

func connectDB(c app.DBConfig) (*gorm.DB, error) {
	dataSource := db.ConnectionString(c.Host, c.Port, c.User, c.Password, c.Name)
	return db.Connect(dataSource)
}

func main() {
	logger := log.NewLogger()
	config, err := app.ReadConfig("app", "./config/")

	if err != nil {
		level.Error(logger).Log("msg",
			fmt.Sprintf("Read config error: %s", err))
		os.Exit(1)
	}

	dbConn, err := connectDB(config.DB)

	if err != nil {
		level.Error(logger).Log("Connect db error: %s", err)
		os.Exit(1)
	}

	sqlDB, err := dbConn.DB()

	if err != nil {
		level.Error(logger).Log("Get sql db from gorm error: %s", err)
		os.Exit(1)
	}

	err = tool.MigrationUp(sqlDB)

	if err != nil {
		level.Error(logger).Log("Migrations error: %s", err)
		os.Exit(1)
	}

	g := &run.Group{}

	http.Start(g, logger, config)
	grpc.Start(g, logger, dbConn, config)
	util.HandleInterrupt(g)

	err = g.Run()

	if err != nil {
		fmt.Println(err)
	}
}
