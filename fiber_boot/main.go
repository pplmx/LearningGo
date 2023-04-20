package main

import (
	"fmt"
	"github.com/pplmx/LearningGo/fiber_boot/bootstrap"
	"github.com/pplmx/LearningGo/fiber_boot/pkg/env"
	"log"
)

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "localhost"), env.GetEnv("APP_PORT", "4000"))))
}
