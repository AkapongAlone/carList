package models

import "gorm.io/gorm"

type Reply struct {
	gorm.Model `json:"-"`
	Text      string `json:"text" gorm:"size:255"`
	CarListID uint `json:"-"`
	CreatedBy uint `json:"-"`
}
