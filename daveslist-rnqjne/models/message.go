package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Message    string
	FromUserID uint
	ToUserID   uint
}
