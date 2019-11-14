package model

import (
	"github.com/jinzhu/gorm"
)

type Bracelet struct {
	gorm.Model
	Slug        string `gorm:"unique_index;not null"`
	text       string `gorm:"not null"`
	threadColor string
	font        string
}

type Bead struct {
	gorm.Model
	Bracelet   Bracelet
	BraceletID uint
	color      string
	design      string
}

