package handler

import "github.com/gofiber/fiber/v2"


func Welcome(c *fiber.Ctx)error {
	data := fiber.Map{
		"message" : "Welcome to soundcloud by valeth",

		"status" : "ok",
	}
	return c.JSON(data)
}

func Get_alltracks(c *fiber.Ctx)error {
	tracks := []fiber.Map{
        {
            "id":     1,
            "title":  "Lagu Keren Pertama",
            "artist": "Valeth",
        },
        {
            "id":     2,
            "title":  "Lagu Keren Kedua",
            "artist": "Sir Valeth",
        },
	}
	return c.JSON(tracks)
}