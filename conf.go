package main

import (
	"fmt"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Account ...
type Account struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func configure(app *echo.Echo, db *mgo.Database) {

	api := app.Group("/api/v1")
	users := db.C("users")

	api.POST("/:user", func(context echo.Context) error {
		account := Account{}
		if err := context.Bind(&account); err != nil ||
			account.Name == "" ||
			account.Link == "" {
			return context.JSON(400, bson.M{
				"error": "You have to provide name, link parameters",
			})
		}
		data := bson.M{}
		if err := users.Find(bson.M{"_username": context.Param("user")}).Select(bson.M{"_id": 0}).One(&data); err != nil {
			err = users.Insert(
				bson.M{
					"_id":        bson.NewObjectId(),
					"_username":  context.Param("user"),
					account.Name: account.Link,
				})
			if err != nil {
				return context.JSON(500, bson.M{
					"error": "Could not create resource",
				})
			}
			return context.JSON(201, data)
		}

		err := users.Update(bson.M{
			"_username": context.Param("user"),
		},
			bson.M{
				"$set": bson.M{account.Name: account.Link},
			})

		if err != nil {
			return context.JSON(500, bson.M{
				"error": "Could not update resource",
			})
		}

		users.Find(bson.M{"_username": context.Param("user")}).Select(bson.M{"_id": 0}).One(&data)
		return context.JSON(201, data)
	})

	api.GET("/:user", func(context echo.Context) error {
		data := bson.M{}
		if err := users.Find(bson.M{"_username": context.Param("user")}).Select(bson.M{"_id": 0}).One(&data); err != nil {
			return context.JSON(404, bson.M{"error": "User does not exist"})
		}
		return context.JSON(200, data)
	})

	api.GET("/:user/:account", func(context echo.Context) error {
		data := bson.M{}
		if err := users.Find(bson.M{"_username": context.Param("user")}).Select(bson.M{"_id": 0, "_username": 1, context.Param("account"): 1}).One(&data); err != nil {
			return context.JSON(404, bson.M{"error": "User does not exist"})
		}
		return context.JSON(200, data)
	})
	app.GET("/", func(context echo.Context) error {
		return context.Redirect(302, "https://github.com/iambenkay/urlify")
	})
	app.GET("/:user/:account", func(context echo.Context) error {
		data := bson.M{}
		if err := users.Find(bson.M{"_username": context.Param("user")}).One(&data); err != nil {
			return context.Redirect(302, "https://github.com/iambenkay/urlify")
		}
		if _, ok := data[context.Param("account")]; !ok {
			return context.Redirect(302, "https://github.com/iambenkay/urlify")
		}

		s := fmt.Sprintf("%v", data[context.Param("account")])
		return context.Redirect(302, s)
	})
}
