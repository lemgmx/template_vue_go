package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"template_api/database"
	"template_api/models"
	"template_api/services"
)

func GetPlanets(c *fiber.Ctx) error {
	var planets []models.Planet
	var count int64

	var offset int
	var page = 1

	queryPage := c.QueryInt("page", 1)
	if queryPage > 1 {
		page = queryPage
	}
	offset = (page - 1) * database.PageLimit

	planets, count = services.GetPlanetsFromDB(offset)

	var next = false
	if int64(offset+database.PageLimit) <= count {
		next = true
	}

	var apiPlanets []models.APIPlanet

	for _, planet := range planets {
		apiPlanets = append(apiPlanets, models.APIPlanet{
			ID:             planet.ID,
			Name:           planet.Name,
			Population:     planet.Population,
			Diameter:       planet.Diameter,
			RotationPeriod: planet.RotationPeriod,
			OrbitalPeriod:  planet.OrbitalPeriod,
		})
	}

	return c.JSON(fiber.Map{
		"count":   count,
		"next":    next,
		"results": apiPlanets,
	})
}

func GetPlanetById(c *fiber.Ctx) error {
	var planet models.Planet
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid planet id",
		})
	}

	planet = services.GetPlanetById(id)
	if planet.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Planet not found",
		})
	}

	var apiPlanet = models.APIPlanet{
		ID:             planet.ID,
		Name:           planet.Name,
		Population:     planet.Population,
		Diameter:       planet.Diameter,
		RotationPeriod: planet.RotationPeriod,
		OrbitalPeriod:  planet.OrbitalPeriod,
	}

	return c.JSON(apiPlanet)
}
