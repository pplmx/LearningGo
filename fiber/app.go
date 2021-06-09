package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "log"
)

func main() {
    app := fiber.New()

    app.Use(cors.New())

    // Or extend your config for customization
    app.Use(cors.New(cors.Config{
        AllowOrigins: "https://gofiber.io, https://gofiber.net",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message": "Hello World 👋!",
        })
    })

    log.Fatal(app.Listen(":3000"))
}
