package main

import (
	"fmt"
	"log"

	"valeth-soundcloud-api/database"
	"valeth-soundcloud-api/routes"
	"valeth-soundcloud-api/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main (){


    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error: Tidak dapat memuat file .env. Pastikan file ada di root.")
    }

	log.Println("initializing databse and storage bucket............")
    database.Init_db()
	storage.InitStorage()
	log.Println("initializing databse and storage bucket succes !")
	//akhirnya bisa


	//kita menset upload limit to 100 MB
	maxUploadSize := 100 * 1024 * 1024

    appConfig := fiber.Config{
        BodyLimit: maxUploadSize,
    }

    app := fiber.New(appConfig)

	routes.Setup_routes(app)
	log.Println("Routes have been assigned")

	//menjalankan server di port 2006
	fmt.Println()
	fmt.Println()
	fmt.Println()
	//listen to port
	var port string = ":2006"
	log.Println("server started at port" +port + "........")
	app.Listen(port)
}

