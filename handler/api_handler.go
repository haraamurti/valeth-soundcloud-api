package handler

import (
	"valeth-soundcloud-api/database"
	"valeth-soundcloud-api/model"

	"github.com/gofiber/fiber/v2"
)



func Welcome(c *fiber.Ctx)error {
	data := fiber.Map{
		"message" : "Welcome to soundcloud by valeth",

		"status" : "ok",
	}
	return c.JSON(data)
}

func Get_alltracks(c *fiber.Ctx)error {
    var tracks []model.Track
	//find dari GOrm
    database.DB.Find(&tracks)
	return c.JSON(tracks)
}