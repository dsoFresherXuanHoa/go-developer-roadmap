package entity

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type Thumb struct {
	gorm.Model
	Url         string                `json:"url" gorm:"column:url;unique;not null"`
	File        *multipart.FileHeader `form:"thumb" sql:"-" gorm:"-"`
	Description string                `json:"description" form:"description" gorm:"column:description;"`
}

type ThumbCreatable struct {
	gorm.Model
	Url         *string               `json:"-" gorm:"column:url;unique;not null"`
	File        *multipart.FileHeader `form:"thumb" sql:"-" gorm:"-"`
	Description *string               `json:"description" form:"description" gorm:"column:description;"`
}

type Thumbs []Thumb

func (Thumb) GetTableName() string          { return "thumbs" }
func (ThumbCreatable) GetTableName() string { return Thumb{}.GetTableName() }
func (Thumbs) GetTableName() string         { return Thumb{}.GetTableName() }
