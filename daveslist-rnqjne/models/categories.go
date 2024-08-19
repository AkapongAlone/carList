package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	CreatedBy    uint
	CategoryName string
	IsPublic     bool
}

type CategoriesPreload struct {
	Categories
	CarList []*CarListPreload `gorm:"ForeignKey:CategoriesID;reference:ID"`
}

func (CategoriesPreload) TableName()string{
	return "categories"
}