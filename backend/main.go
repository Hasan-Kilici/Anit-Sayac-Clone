package main

import (
	"anitsayac/middlewares"
	"anitsayac/routers"
	"anitsayac/cronjobs"

	"github.com/gofiber/fiber/v2"
	"github.com/goccy/go-json"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(middlewares.Cors)
	app.Use(middlewares.Logger)
	app.Use(middlewares.Compress)
	app.Use(middlewares.Security)
	app.Use(middlewares.RateLimit)

	api := app.Group("/api")
	routers.Api(api)

	app.Use(middlewares.NotFound)

	cronScheduler := cronjobs.InitializeCron()

	err := app.Listen("127.0.0.1:4000")
	if err != nil {
		panic(err)
	}

	defer cronScheduler.Stop()

	select {}
}
