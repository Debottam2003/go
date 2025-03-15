package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var port string = ":5000"
	app := fiber.New()
	fmt.Println("Server is running on port: ", port)
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Request came successfully in my first go backend...")
		return c.SendString("hello world")
	})
	app.Listen(port)
}
