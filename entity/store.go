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

type StoreWithDorayakiAmount struct {
	Store
	DorayakisAmount []DorayakiWithAmount
}

func (store Store) Transform() StoreWithDorayakiAmount {
	var result StoreWithDorayakiAmount
	result.Model = store.Model
	result.Address = store.Address
	result.Name = store.Name
	result.Kecamatan = store.Kecamatan
	result.Province = store.Province
	result.Dorayakis = store.Dorayakis
	result.DorayakisAmount = make([]DorayakiWithAmount, 0)
	return result
}

func (Store) TableName() string {
	return "store"
}
