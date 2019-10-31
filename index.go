package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	app := echo.New()
	app.Static("/static", "assets")
	app.GET("/favicon.ico", func(context echo.Context) error {
		return context.File("./assets/res/favicon.ico")
	})
	configure(app, session.DB("urlify"))

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method} uri=${uri} status=${status}\n",
	}))
	app.Use(middleware.Recover())

	app.Logger.Fatal(app.Start(":8000"))
}
