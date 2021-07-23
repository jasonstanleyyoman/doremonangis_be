package repo_gorm

import (
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
	repo.Db.Find(&dorayaki, id)
	return dorayaki, nil
}

func (repo *DorayakiRepo) AddDorayaki(dorayaki *entity.Dorayaki) {
	repo.Db.Create(dorayaki)
}

func (repo *DorayakiRepo) RemoveDorayaki(id uint) {
	repo.Db.Where("dorayaki_id = ?", id).Delete(&entity.StoreDorayaki{})
	repo.Db.Unscoped().Where("id = ?", id).Delete(&entity.Dorayaki{})
}
