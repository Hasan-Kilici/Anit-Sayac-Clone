package routers

import (
	"anitsayac/handlers"
	"github.com/gofiber/fiber/v2"
)

func Api(app fiber.Router) {
	app.Get("/list/incidents", handlers.GetIncidents)
	app.Get("/search/incidents", handlers.SearchIncidents)
	app.Get("/find/incident", handlers.FindIncident)
}