package routes

import (
    "LearningGo/fiber/handlers"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func New() *fiber.App {
    app := fiber.New()
    app.Use(cors.New())
    app.Use(logger.New(logger.Config{
        Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${magenta}${path}\n",
        TimeFormat: "2006-01-02 15:04:05",
        TimeZone:   "Asia/Shanghai",
    }))

    api := app.Group("/api") // /api

    api.Get("/hello", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message": "Hello World 👋!",
        })
    })

    v1 := api.Group("/v1") // /api/v1
    userRouter := v1.Group("users")
    userRouter.Get("/", handlers.GetAllUsers)           // /api/v1/users/
    userRouter.Get("/:uid", handlers.GetUserByID)       // /api/v1/users/:uid
    userRouter.Delete("/:uid", handlers.DeleteUserByID) // /api/v1/users/:uid

    v2 := api.Group("/v2")                  // /api/v2
    v2.Get("/list", handlers.GetAllUsers)   // /api/v2/list
    v2.Get("/user", handlers.GetUserByName) // /api/v2/user

    return app
}

func HttpNotFound(c *fiber.Ctx) error {
    return c.Status(404).SendFile("fiber/static/404.html")
}
