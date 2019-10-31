package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
)

func main() {
	var mongo string
	var dbName string

	if mongo = os.Getenv("MONGODB_URI"); mongo == "" {
		mongo = "localhost"
	}
	if dbName = os.Getenv("DB_NAME"); dbName == "" {
		dbName = "urlspace"
	}

	session, err := mgo.Dial(mongo)

	if err != nil {
		panic(err)
	}

	app := echo.New()

	configure(app, session.DB(dbName))

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method} uri=${uri} status=${status}\n",
	}))
	app.Use(middleware.Recover())

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", port)))
}
