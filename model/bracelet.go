package model

import (
	"github.com/jinzhu/gorm"
)

type Bracelet struct {
	gorm.Model
	Slug        string `gorm:"unique_index;not null"`
	Text       string `gorm:"not null"`
	ThreadColor string
	Font        string
}

type Bead struct {
	gorm.Model
	Bracelet   Bracelet
	BraceletID uint
	Color      string
	Design      string
}

