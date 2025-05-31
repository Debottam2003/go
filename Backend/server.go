package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string
	Age  int
}

func main() {
	var port string = ":5000"
	//app := fiber.New()
	app := fiber.New(fiber.Config{Prefork: true})
	fmt.Println("Server is running on port", port)
	app.Get("/", func(c *fiber.Ctx) error {
		//fmt.Println("Request came successfully in my first go backend...")
		user := User{Name: "Debottam Kar", Age: 22}
		// return c.SendString("Hello World!")
		return c.JSON(user)
	})
	app.Listen(port)
}
