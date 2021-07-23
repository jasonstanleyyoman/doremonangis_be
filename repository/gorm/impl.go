package repo_gorm

import (
	"github.com/jasonstanleyyoman/doremonangis_be/entity"
	repository_base "github.com/jasonstanleyyoman/doremonangis_be/repository"
	"gorm.io/gorm"
)

type GormRepo struct {
	Db           gorm.DB
	StoreRepo    repository_base.IStoreRepo
	DorayakiRepo repository_base.IDorayakiRepo
}

func (repo *GormRepo) InitRepo() {
	repo.Db.AutoMigrate(&entity.Dorayaki{})
	repo.Db.AutoMigrate(&entity.Store{})
	repo.Db.AutoMigrate(&entity.StoreDorayaki{})
	repo.Db.SetupJoinTable(&entity.Store{}, "dorayaki", &entity.StoreDorayaki{})

	repo.StoreRepo = &StoreRepo{
		Db: repo.Db,
	}
	repo.DorayakiRepo = &DorayakiRepo{
		Db: repo.Db,
	}
}

func (repo *GormRepo) GetStoreRepo() repository_base.IStoreRepo {
	return repo.StoreRepo
}

func (repo *GormRepo) GetDorayakiRepo() repository_base.IDorayakiRepo {
	return repo.DorayakiRepo
}
