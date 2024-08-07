package main

import (
	"fmt"
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
	redis := database.InitRedis(cfg)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	route.New(e, dbMysql, supabase, redis)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	port := cfg.SERVERPORT
	if port == 0 {
		port = 8000
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
