package routes

import (
	"github.com/gofiber/fiber/v2"
	"template_api/controllers"
)

func PlanetsRoutes(app *fiber.App) {
	app.Get("/planets", controllers.GetPlanets)
	app.Get("/planets/:id?", controllers.GetPlanetById)
}
