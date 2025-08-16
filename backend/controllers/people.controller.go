package controllers

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"template_api/database"
	"template_api/models"
	"template_api/services"

	"github.com/gofiber/fiber/v2"
)

func GetPeople(c *fiber.Ctx) error {
	var people []models.Person
	var count int64

	var offset int
	var page = 1

	queryPage := c.QueryInt("page", 1)
	if queryPage > 1 {
		page = queryPage
	}
	offset = (page - 1) * database.PageLimit

	people, count = services.GetPeopleFromDB(offset)

	var next = false
	if int64(offset+database.PageLimit) <= count {
		next = true
	}

	var apiPeople []models.APIPerson

	for _, person := range people {
		apiPeople = append(apiPeople, models.APIPerson{
			ID:        person.ID,
			Name:      person.Name,
			Gender:    person.Gender,
			Height:    person.Height,
			Mass:      person.Mass,
			HairColor: person.HairColor,
			Homeworld: os.Getenv("APP_DOMAIN") + "/planets/" + strconv.Itoa(person.PlanetID),
			URL:       strings.Replace(person.URL, "https://swapi.dev/api", os.Getenv("APP_DOMAIN"), -1),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"count":   count,
		"next":    next,
		"results": apiPeople,
	})
}

func GetPersonById(c *fiber.Ctx) error {
	var person models.Person
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid planet id",
		})
	}

	person = services.GetPersonById(id)
	if person.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Person not found",
		})
	}

	var apiPerson = models.APIPerson{
		ID:        person.ID,
		Name:      person.Name,
		Gender:    person.Gender,
		Height:    person.Height,
		Mass:      person.Mass,
		HairColor: person.HairColor,
		Homeworld: os.Getenv("APP_DOMAIN") + "/planets/" + strconv.Itoa(person.PlanetID),
		URL:       strings.Replace(person.URL, "https://swapi.dev/api", os.Getenv("APP_DOMAIN"), -1),
	}

	return c.JSON(apiPerson)
}
