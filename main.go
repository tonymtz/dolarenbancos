package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	v1 "github.com/tonymtz/dolarenbancos/api/v1"
	"github.com/tonymtz/dolarenbancos/scheduler"
	"github.com/tonymtz/dolarenbancos/server/logger"
	"os"
)

func main() {
	db, err := sql.Open("postgres", getDbString())

	if err != nil {
		logger.Error(err)
	}

	scheduler.StartScheduler(db)

	e := echo.New()

	v1.Router(e)

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	e.Logger.Fatal(e.Start(address))
}

func getDbString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		host,
		port,
		user,
		name,
		password,
	)
}
