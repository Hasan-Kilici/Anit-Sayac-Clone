package handlers

import (
	"fmt"
	"anitsayac/database"
	"github.com/gofiber/fiber/v2"
)

func GetIncidents(c *fiber.Ctx) error {
	year := c.Query("year")
	if(year != ""){
		incidents, err := database.ListIncidentsByYear(year)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("%s", err))
		}
	
		return c.JSON(fiber.Map{
			"data": incidents,
		})
	} else {
		incidents, err := database.ListIncidents()
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("%s", err))
		}

		return c.JSON(fiber.Map{
			"data": incidents,
		})
	}
}

func SearchIncidents(c *fiber.Ctx) error {
	name := c.Query("name")
	incidents, err := database.SearchIncidentsByName(name)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("%s", err))
	}

	return c.JSON(fiber.Map{
		"data": incidents,
	})
}

func FindIncident(c *fiber.Ctx) error {
	id := c.QueryInt("id");
	incident, err := database.GetIncidentById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("%s", err))
	} 

	return c.JSON(fiber.Map{
		"data": incident,
	})
}
