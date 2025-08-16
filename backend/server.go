package main

import (
	"os"
	"template_api/database"
	"template_api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	if os.Getenv("APP_DEBUG") == "1" {
		app.Use(logger.New())
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("APP_ORIGINS") + "," + os.Getenv("APP_DOMAIN"),
		AllowMethods:     "GET,POST",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	err = database.Connect()
	if err != nil {
		log.Error("Unable to connect to database")
		return
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.PeopleRoutes(app)
	routes.PlanetsRoutes(app)

	err = app.Listen(":3000")

	if err != nil {
		log.Error("Unable to start app ", err)
		return
	}
}
