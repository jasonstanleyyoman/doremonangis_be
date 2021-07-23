package entity

import (
	"gorm.io/gorm"
)

type Dorayaki struct {
	gorm.Model
	Flavor      string
	Description string
	ImagePath   string
	Stores      []*Store `gorm:"many2many:store_dorayaki"`
}

func (Dorayaki) TableName() string {
	return "dorayaki"
}
