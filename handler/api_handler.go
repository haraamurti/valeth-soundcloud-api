package handler

import (
	"fmt"
	"log"
	"path/filepath"
	"valeth-soundcloud-api/database"
	"valeth-soundcloud-api/model"
	"valeth-soundcloud-api/storage"

	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func CreateTrack(c *fiber.Ctx)error{
	log.Println("Handler : accepting request CreateTrack")

	title := c.FormValue("title")
	artist := c.FormValue("artist")

	if (title == "" || artist == ""){
		log.Println("fill the title and artist")

		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"message": "fill the title and artist"})
	}

	//track file audio get
	trackHeader, err := c.FormFile("track_file")

    if err != nil {
        log.Println("Handler Error: track file is missing")
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "fill the track file"})
    }

	//cover file selection get
    coverHeader, err := c.FormFile("cover_file")
    if err != nil {
        log.Println("Handler Error: cover file is missing")
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "fill the cover file"})
    }

	//uploading audio file audio selection
    log.Println("Handler: Memproses file audio...")
    trackFile, err := trackHeader.Open()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "error opening file audio"})
    }
	defer trackFile.Close()
	

	trackExt := filepath.Ext(trackHeader.Filename)
	trackName := fmt.Sprintf("track-%s%s", uuid.New().String(), trackExt)

	trackURL, err := storage.UploadFile("tracks", trackName, trackFile, trackHeader.Header.Get("Content-Type"))
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "audio upload failed", "detail": err.Error()})
    }
	log.Println("audio upload succesfull!")

	log.Println("handler : processing file cover.......")
	coverFile, err := coverHeader.Open()
	if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "error opening file cover"})
    }
	defer trackFile.Close()
	

	coverExt := filepath.Ext(coverHeader.Filename)
    coverName := fmt.Sprintf("cover-%s%s", uuid.New().String(), coverExt)

	coverURL, err := storage.UploadFile("tracks", coverName, coverFile, coverHeader.Header.Get("Content-Type"))
	if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "cover image upload failed", "detail": err.Error()})
    }
	log.Println("cover image upload succesfull!")
	
	log.Println("handle : saving track metadata to database ")

	track := model.Track{
    Title:         title,
    Artist:        artist,
    TrackURL:      trackURL,
    TrackCoverURL: coverURL,
}

	if err := database.DB.Create(&track).Error; err != nil{
		log.Println("Handler error : failed to save to DB ", err.Error())
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "track metadata save failed", "detail": err.Error()})
	}
	log.Println("create track succes")
	return c.Status(201).JSON(fiber.Map{
		"status" : "succes",
		"message" : "Track uploaded succesfully",
		"data" : track,
	})
}


func Get_track_by_id (c *fiber.Ctx)error{
	id := c.Params("id")
	var track model.Track
	result:=database.DB.First(&track, id)

	if result.Error != nil{
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return c.Status(404).JSON(fiber.Map{
				"status": "Error",
				"message": "file not found in database",
			})
		}
		return c.Status(500).JSON(fiber.Map{
				"status": "Error",
				"message": "internal server error",
			})
		}
		return c.JSON(track)
	}
//retrieve audio file

func Get_track_audio(c *fiber.Ctx)error{
		id := c.Params("id")
		var track model.Track
		result:=database.DB.First(&track, id)

		if result.Error != nil{
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return c.Status(404).JSON(fiber.Map{
				"status": "Error",
				"message": "file not found in database",
			})
		}
		return c.Status(500).JSON(fiber.Map{
				"status": "Error",
				"message": "internal server error",
			})
		}

		return c.Redirect(track.TrackURL, fiber.StatusFound)

		}

//retrive cover url

func Get_track_cover(c *fiber.Ctx)error{
		id := c.Params("id")
		var track model.Track
		result:=database.DB.First(&track, id)

		if result.Error != nil{
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return c.Status(404).JSON(fiber.Map{
				"status": "Error",
				"message": "file not found in database",
			})
		}
		return c.Status(500).JSON(fiber.Map{
				"status": "Error",
				"message": "internal server error",
			})
		}

		return c.Redirect(track.TrackCoverURL, fiber.StatusFound)
		}


func Edit_title_and_artist (c *fiber.Ctx)error{
	id := c.Params(":id")
	var track model.Track
	result := database.DB.First(&track, id)
	if result.Error != nil{
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return c.Status(404).JSON(fiber.Map{
				"status": "Error",
				"message": "file not found in database",
			})
		}
		return c.Status(500).JSON(fiber.Map{
				"status": "Error",
				"message": "internal server error",
			})
		}
		var input map[string]interface{}
		if err := c.BodyParser(&input); err != nil {
    	return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"message": "Input JSON tidak valid"})
		}
		if err := database.DB.Model(&track).Updates(input).Error; err != nil {
    	return c.Status(500).JSON(fiber.Map{
				"status": "error",
				"message": "Gagal update data"})
}
	return c.JSON(track)
}

func Delete(c *fiber.Ctx)error{
	id := c.Params(":id")
	var track model.Track
	result := database.DB.First(&track, id)
	if result.Error != nil{
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return c.Status(404).JSON(fiber.Map{
				"status": "Error",
				"message": "file not found in database",
			})
		}
		return c.Status(500).JSON(fiber.Map{
				"status": "Error",
				"message": "internal server error",
			})
		}
		if err := database.DB.Delete(&track).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"status": "Error",
				"message": "canot delete file",
			})
		}
		return c.Status(204).JSON(track)
}
