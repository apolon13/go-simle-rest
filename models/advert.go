package models

import "gorm.io/gorm"

type Advert struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}
