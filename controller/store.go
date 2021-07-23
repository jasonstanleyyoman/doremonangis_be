package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonstanleyyoman/doremonangis_be/connection"
	repository_base "github.com/jasonstanleyyoman/doremonangis_be/repository"
	repo_gorm "github.com/jasonstanleyyoman/doremonangis_be/repository/gorm"
)

type StoreController struct {
	Repo repository_base.MasterRepo
}

func NewStoreController() StoreController {
	repo := repo_gorm.GormRepo{
		Db: connection.GetConnection().Db,
	}
	repo.InitRepo()
	return StoreController{
		Repo: &repo,
	}
}

func (cont *StoreController) GetAllStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		stores := cont.Repo.GetStoreRepo().GetAllStore()
		c.JSON(http.StatusOK, stores)
	}
}
