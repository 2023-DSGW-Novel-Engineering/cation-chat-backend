package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `json:"name"`
	UserID         string `gorm:"url" json:"user_id"`
	Password       string `json:"password"`
	NativeLanguage string `json:"native_language"`
}
