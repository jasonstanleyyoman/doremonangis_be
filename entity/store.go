package entity

import (
	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	Name      string
	Address   string
	Kecamatan string
	Province  string
	Dorayakis []*Dorayaki `gorm:"many2many:store_dorayaki"`
}

func (Store) TableName() string {
	return "store"
}
