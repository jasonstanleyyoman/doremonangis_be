package repository_base

import (
	"github.com/jasonstanleyyoman/doremonangis_be/entity"
)

type IDorayakiRepo interface {
	GetAllDorayaki() []entity.Dorayaki
	GetDorayakiInfo(id uint) (entity.Dorayaki, error)
	AddDorayaki(dorayaki entity.Dorayaki) (entity.Dorayaki, error)
	RemoveDorayaki(id uint) error
}

type IStoreRepo interface {
	GetAllStore() []entity.StoreWithDorayakiAmount
	GetStoreInfo(id uint) (entity.StoreWithDorayakiAmount, error)
	AddStore(store entity.Store) (entity.Store, error)
	RemoveStore(id uint) error
	AddStock(storeId, dorayakiId uint, amount int) error
	RemoveStock(storeId, dorayakiId uint, amount int) error
	MoveStock(srcId, destId, dorayakiId uint, amount int) error
	CheckStore(storeId uint) bool
}

type MasterRepo interface {
	GetStoreRepo() IStoreRepo
	GetDorayakiRepo() IDorayakiRepo
	InitRepo()
}
