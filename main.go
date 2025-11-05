package main

import (
	"fmt"
	"log"

	"valeth-soundcloud-api/database"
	"valeth-soundcloud-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main (){
	err :=godotenv.Load()
	if err != nil {
		log.Fatal("Error: cannot find file .env")
	}
	app := fiber.New()
	database.Init_db()
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

