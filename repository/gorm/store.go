package repo_gorm

import (
	"errors"

	"github.com/jasonstanleyyoman/doremonangis_be/entity"
	"gorm.io/gorm"
)

type StoreRepo struct {
	Db gorm.DB
}

func (repo *StoreRepo) GetAllStore() []entity.Store {
	var stores []entity.Store
	repo.Db.Preload("Dorayakis").Find(&stores)
	return stores
}

func (repo *StoreRepo) GetStoreInfo(id uint) (entity.Store, error) {
	var store entity.Store
	repo.Db.Preload("Dorayakis").Find(&store, id)
	return store, nil
}

func (repo *StoreRepo) AddStore(store *entity.Store) error {
	err := repo.Db.Create(&store)

	return err.Error
}

func (repo *StoreRepo) RemoveStore(id uint) error {
	repo.Db.Where("store_id = ?", id).Delete(&entity.StoreDorayaki{})
	repo.Db.Unscoped().Delete(&entity.Store{}, id)
	return nil
}

func (repo *StoreRepo) AddStock(storeId, dorayakiId uint, amount int) error {

	if err := repo.Db.Where("store_id = ? AND dorayaki_id = ?", storeId, dorayakiId).First(&entity.StoreDorayaki{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		storeDorayaki := &entity.StoreDorayaki{
			StoreId:    storeId,
			DorayakiId: dorayakiId,
			Amount:     amount,
		}
		repo.Db.Create(storeDorayaki)
	} else {
		repo.Db.Model(&entity.StoreDorayaki{}).Where("store_id = ? AND dorayaki_id = ?", storeId, dorayakiId).Update("amount", gorm.Expr("amount + ?", amount))

	}
	return nil
}

func (repo *StoreRepo) RemoveStock(storeId, dorayakiId uint, amount int) error {
	var storeDorayaki entity.StoreDorayaki

	repo.Db.Where("store_id = ? AND dorayaki_id = ?", storeId, dorayakiId).First(&storeDorayaki)
	storeDorayaki.Amount -= amount
	repo.Db.Save(&storeDorayaki)

	return nil
}
