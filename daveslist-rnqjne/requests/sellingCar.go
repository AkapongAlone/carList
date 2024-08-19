package requests

import "mime/multipart"

type RequestSellingCar struct {
	CategoryID uint                    `form:"category_id" binding:"required"`
	Title      string                  `form:"title" binding:"required"`
	Content    string                  `form:"content" binding:"required"`
	File       []*multipart.FileHeader `form:"file" binding:"required"`
	IsPublic   bool                    `form:"is_public" binding:"required"`
}

type RequestHideList struct {
	IsShow     bool                    `json:"is_show"`
}


type RequestCategory struct {
	CategoryName string `json:"category_name"  binding:"required"`
	IsPublic     bool   `json:"is_public"  binding:"required"`
}

type Reply struct {
	CarListID uint `json:"-"  `
	Text      string `json:"text"  binding:"required"`
}
