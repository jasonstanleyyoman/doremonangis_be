package repo_gorm

import (
	"errors"

	"github.com/jasonstanleyyoman/doremonangis_be/entity"
	"gorm.io/gorm"
)

type DorayakiRepo struct {
	Db gorm.DB
}

func (repo *DorayakiRepo) GetAllDorayaki() []entity.Dorayaki {
	var dorayakis []entity.Dorayaki
	repo.Db.Find(&dorayakis)
	return dorayakis
}

func (repo *DorayakiRepo) GetDorayakiInfo(id uint) (entity.Dorayaki, error) {
	var dorayaki entity.Dorayaki
	if err := repo.Db.First(&dorayaki, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return dorayaki, err
	}
	return dorayaki, nil
}

func (repo *DorayakiRepo) AddDorayaki(dorayaki entity.Dorayaki) (entity.Dorayaki, error) {
	err := repo.Db.Create(&dorayaki)

	return dorayaki, err.Error
}

func (repo *DorayakiRepo) RemoveDorayaki(id uint) error {
	if err := repo.Db.Where("id = ?", id).First(&entity.Dorayaki{}).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	repo.Db.Where("dorayaki_id = ?", id).Delete(&entity.StoreDorayaki{})
	repo.Db.Unscoped().Where("id = ?", id).Delete(&entity.Dorayaki{})

	return nil
}
