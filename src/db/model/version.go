package model

import "gorm.io/gorm"

type FileVersion struct {
	gorm.Model
	Version int
	FileID
}
