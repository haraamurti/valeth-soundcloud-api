package handler

import (
	"valeth-soundcloud-api/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB*gorm.DB

func Welcome(c *fiber.Ctx)error {
	data := fiber.Map{
		"message" : "Welcome to soundcloud by valeth",

		"status" : "ok",
	}
	return c.JSON(data)
}

func Get_alltracks(c *fiber.Ctx)error {
    var tracks []model.Track
    DB.Find(&tracks)
	return c.JSON(tracks)
}