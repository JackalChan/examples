package main

import (
	"time"

	"github.com/iris-contrib/graceful"
	"gopkg.in/kataras/iris.v6"
)

func main() {
	api := iris.New()
	api.Get("/", func(c *iris.Context) {
		c.Writef("Welcome to the home page!")
	})

	graceful.Run(":3001", time.Duration(20)*time.Second, api)
}
