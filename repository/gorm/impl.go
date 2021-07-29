package repo_gorm

import (
	"sync"

	"github.com/jasonstanleyyoman/doremonangis_be/connection"
	"github.com/jasonstanleyyoman/doremonangis_be/entity"
	repository_base "github.com/jasonstanleyyoman/doremonangis_be/repository"
	"gorm.io/gorm"
)

var (
	mut  sync.Mutex = sync.Mutex{}
	repo *GormRepo
)

type GormRepo struct {
	Db           gorm.DB
	StoreRepo    repository_base.IStoreRepo
	DorayakiRepo repository_base.IDorayakiRepo
	initialized  bool
}

func (repo *GormRepo) InitRepo() {
	if repo.initialized {
		return
	}

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
	repo.initialized = true
}

func (repo *GormRepo) GetStoreRepo() repository_base.IStoreRepo {
	return repo.StoreRepo
}

func (repo *GormRepo) GetDorayakiRepo() repository_base.IDorayakiRepo {
	return repo.DorayakiRepo
}

func GetRepo() repository_base.MasterRepo {
	if repo != nil && repo.initialized {
		return repo
	}
	mut.Lock()
	defer mut.Unlock()
	if repo == nil {

		repo = &GormRepo{
			Db: connection.GetConnection().Db,
		}
		repo.InitRepo()
	}
	return repo
}
