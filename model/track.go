package model

import "gorm.io/gorm"

type Track struct{

	gorm.Model
    Title  string `json:"title"`
    Artist string `json:"artist"`
	TrackURL string `json:"track_url"`
	TrackCoverURL string `json:"cover_url"`

}