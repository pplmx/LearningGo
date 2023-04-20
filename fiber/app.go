package main

import (
	"github.com/pplmx/LearningGo/fiber/routes"
	"log"
)

func main() {
	app := routes.New()

	// Handler 404
	app.Use(routes.HttpNotFound)

	log.Fatal(app.Listen(":3000"))
}
