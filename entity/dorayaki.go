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

type DorayakiWithAmount struct {
	Dorayaki
	Amount int
}

func (dorayaki Dorayaki) Transform() DorayakiWithAmount {
	var result DorayakiWithAmount
	result.Flavor = dorayaki.Flavor
	result.Model = dorayaki.Model
	result.Description = dorayaki.Description
	result.ImagePath = dorayaki.ImagePath
	result.Amount = 0
	return result
}

func (Dorayaki) TableName() string {
	return "dorayaki"
}
