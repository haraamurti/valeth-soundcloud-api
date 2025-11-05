package main

import (
	"fmt"
	"log"

	"valeth-soundcloud-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main (){
	app := fiber.New()
	routes.Setup_routes(app)

	//menjalankan server di port 2006
	fmt.Println()
	fmt.Println()
	fmt.Println()
	//listen to port
	var port string = ":2006"
	log.Println("server started at port" +port + "........")
	app.Listen(port)
}

