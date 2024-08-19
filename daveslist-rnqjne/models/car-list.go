package models

import (
	"gorm.io/gorm"
)

type CarList struct {
	gorm.Model   `json:"-"`
	CreatedBy    uint   `json:"-"`
	CategoriesID uint   `json:"-"`
	Title        string `json:"title"`
	Content      string `json:"content" gorm:"size:5000"`
	IsPublic     bool   `json:"-"`
	IsShow       bool   `json:"-"`
	IsTrash      bool   `json:"-"`
}

type CarListPreload struct {
	CarList
	Files   []*File  `gorm:"ForeignKey:CarListID;reference:ID"`
	Replies []*Reply `gorm:"ForeignKey:CarListID;reference:ID"`
}

func (CarListPreload) TableName() string {
	return "car_lists"
}

