package routes

import (
	"valeth-soundcloud-api/handler"

	"github.com/gofiber/fiber/v2"
)

func Setup_routes (app *fiber.App){

	//get routes
	app.Get("/",handler.Welcome)
	app.Get("/tracks",handler.Get_alltracks)
	app.Get("/tracks/:id",handler.Get_track_by_id)
	app.Get("/tracks/:id/audio",handler.Get_track_audio)
	app.Get("/tracks/:id/cover",handler.Get_track_cover)
	
	//post routes
	app.Post("/upload-track",handler.CreateTrack)


	app.Patch("/tracks/:id/edit",handler.Edit_title_and_artist)

	
}
