package models

import (
	"time"
)

// A struct to hold music info
type Music struct {
	Music_id      int
	Full_name     string
	Album         string
	Genre         string
	Date_released time.Time
	Artist        string
}

//A struct to hold a user
