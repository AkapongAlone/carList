package models

import "gorm.io/gorm"

type File struct {
	gorm.Model `json:"-"`
	URL      string `json:"url"`
	CarListID uint `json:"-"`
}
