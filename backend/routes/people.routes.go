package routes

import (
	"github.com/gofiber/fiber/v2"
	"template_api/controllers"
)

func PeopleRoutes(app *fiber.App) {
	app.Get("/people", controllers.GetPeople)
	app.Get("/people/:id?", controllers.GetPersonById)
}
