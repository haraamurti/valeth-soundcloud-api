package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main (){
	app := fiber.New()

	app.Get("/", func (c *fiber.Ctx) error {
		return c.JSON("Welcome to soundcloud by valeth")
	})

	//menjalankan server di port 2006
	fmt.Println()
	fmt.Println()
	fmt.Println()
	log.Println("server started....")
	app.Listen(":2006")
}

