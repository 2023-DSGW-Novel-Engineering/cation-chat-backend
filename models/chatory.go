package models

import "gorm.io/gorm"

type Chatory struct {
	gorm.Model
	// UserRealID uint
	Name    string
	Content string
}
