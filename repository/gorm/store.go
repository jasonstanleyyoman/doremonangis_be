package repo_gorm

import (
	"errors"

	"github.com/jasonstanleyyoman/doremonangis_be/entity"
	"gorm.io/gorm"
)

type StoreRepo struct {
	Db gorm.DB
}

func (repo *StoreRepo) GetAllStore() []entity.StoreWithDorayakiAmount {
	var stores []entity.Store
	results := make([]entity.StoreWithDorayakiAmount, 0)
	repo.Db.Preload("Dorayakis").Find(&stores)

	for _, store := range stores {
		storeInfo, _ := repo.GetStoreInfo(store.ID)
		results = append(results, storeInfo)
	}
	return results
}

func (repo *StoreRepo) GetStoreInfo(id uint) (entity.StoreWithDorayakiAmount, error) {
	var store entity.Store
	if tx := repo.Db.Preload("Dorayakis").First(&store, id); tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return entity.StoreWithDorayakiAmount{}, tx.Error
	}
	result := store.Transform()

	for _, dorayaki := range result.Dorayakis {
		dorayakiAmount := dorayaki.Transform()
		var storeDorayaki entity.StoreDorayaki
		repo.Db.Where("store_id = ? AND dorayaki_id = ?", store.ID, dorayaki.ID).First(&storeDorayaki)
		dorayakiAmount.Amount = storeDorayaki.Amount
		result.DorayakisAmount = append(result.DorayakisAmount, dorayakiAmount)
	}
	return result, nil
}

func (repo *StoreRepo) AddStore(store entity.Store) (entity.Store, error) {
	err := repo.Db.Create(&store)

	return store, err.Error
}

func (repo *StoreRepo) RemoveStore(id uint) error {
	if err := repo.Db.Debug().Where("id = ?", id).First(&entity.Store{}).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	repo.Db.Where("store_id = ?", id).Delete(&entity.StoreDorayaki{})
	repo.Db.Unscoped().Delete(&entity.Store{}, id)
	return nil
}

func (repo *StoreRepo) AddStock(storeId, dorayakiId uint, amount int) error {
	if !repo.CheckStore(storeId) {
		return gorm.ErrRecordNotFound
	}
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
	if !repo.CheckStore(storeId) {
		return gorm.ErrRecordNotFound
	}
	var storeDorayaki entity.StoreDorayaki

	if err := repo.Db.Where("store_id = ? AND dorayaki_id = ?", storeId, dorayakiId).First(&storeDorayaki).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	storeDorayaki.Amount -= amount
	repo.Db.Save(&storeDorayaki)

	return nil
}

func (repo *StoreRepo) MoveStock(srcId, destId, dorayakiId uint, amount int) error {
	var sourceStockInfo entity.StoreDorayaki
	if err := repo.Db.First(&sourceStockInfo, srcId).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("source store doesnt have the dorayaki")
	}
	if sourceStockInfo.Amount < amount {
		return errors.New("not enough dorayaki to be moved")
	}
	if !(repo.CheckStore(srcId) && repo.CheckStore(destId)) {
		return errors.New("store not found")
	}
	repo.RemoveStock(srcId, dorayakiId, amount)
	repo.AddStock(destId, dorayakiId, amount)
	return nil
}

func (cont *StoreRepo) CheckStore(storeId uint) bool {
	if err := repo.Db.First(&entity.Store{}, storeId).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
