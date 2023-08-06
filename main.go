package main

import (
	"log"

	c "github.com/ashpect/obs-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	api := app.Group("/api/v1")

	admin := api.Group("/admin")

	admin.Get("/hitobsapi", c.CallBasicApi)
	admin.Get("/hitobsapi_creds", c.CredsHitApi)

	log.Fatal(app.Listen(":3000"))
}
