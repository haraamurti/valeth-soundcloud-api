package database

import (
	"log"
	"os"
	"valeth-soundcloud-api/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB*gorm.DB

func Init_db(){
dsn := os.Getenv("DB_URL")

if dsn == ""{
	log.Fatal("Error: Variabel DB_URL does not exsit  in.env")
}
d, err :=gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
	log.Fatal("Connection failed")
}

log.Println("database connected!")

d.AutoMigrate(&model.Track{})
log.Println("Database migrated!")

DB = d
}