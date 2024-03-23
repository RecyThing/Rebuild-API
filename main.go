package main

import (
	"fmt"
	"log"
	"recything/app/config"
	"recything/app/database"
	"recything/app/route"
	"recything/utils/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	cfg := config.InitConfig()
	dbMysql := database.InitDBMysql(cfg)
	database.InitMigrationMysql(dbMysql)
	supabase := storage.InitStorage(cfg)
	log.Print(cfg.API_STORAGE)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	route.New(e, dbMysql,supabase)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVERPORT)))
}
