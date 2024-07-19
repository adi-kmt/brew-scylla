package main

import (
	"fmt"
	"log"

	"github.com/adi-kmt/brew-scylla/internal/authflow"
	"github.com/adi-kmt/brew-scylla/internal/common/injector"
	user_controllers "github.com/adi-kmt/brew-scylla/internal/user/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	deps, err := injector.Inject()
	if err != nil {
		panic(err)
	}

	api := app.Group("/api/v1")

	app.Get("/metrics", monitor.New(monitor.Config{Title: "BrewScylla Metrics Page"}))

	defer deps.Session.Close()

	user_controllers.UserHandler(api, deps.Order, deps.Product)
	authflow.AuthHandler(api, deps.User)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", 8080)))
}
