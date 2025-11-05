package model

import "gorm.io/gorm"

type Track struct{

	gorm.Model
    Title  string `json:"title"`
    Artist string `json:"artist"`
	TrackUrl string `json:"track_url"`
	TrackCoverUrl string `json:"cover_url"`

}