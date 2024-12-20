package main

import (
	"log"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/Web-developing-team/user_management_module/user_management_model"
	"github.com/Web-developing-team/user_management_module/user_management_routes"
)

func main() {

	fmt.Println(user_management_model.User{})

	// Initialize Fiber app
	app := fiber.New()

	user_management_routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}