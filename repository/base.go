package repository_base

import (
	"github.com/jasonstanleyyoman/doremonangis_be/entity"
)

type IDorayakiRepo interface {
	GetAllDorayaki() []entity.Dorayaki
	GetDorayakiInfo(id uint) (entity.Dorayaki, error)
	AddDorayaki(dorayaki *entity.Dorayaki)
	RemoveDorayaki(id uint)
}

type IStoreRepo interface {
	GetAllStore() []entity.Store
	GetStoreInfo(id uint) (entity.Store, error)
	AddStore(store *entity.Store) error
	RemoveStore(id uint) error
	AddStock(storeId, dorayakiId uint, amount int) error
	RemoveStock(storeId, dorayakiId uint, amount int) error
}

type MasterRepo interface {
	GetStoreRepo() IStoreRepo
	GetDorayakiRepo() IDorayakiRepo
	InitRepo()
}
