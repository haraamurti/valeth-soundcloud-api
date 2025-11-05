package routes

import (
	"valeth-soundcloud-api/handler"

	"github.com/gofiber/fiber/v2"
)

func Setup_routes (app *fiber.App){
	app.Get("/",handler.Welcome)
	
	app.Get("/tracks",handler.Get_alltracks)
}
