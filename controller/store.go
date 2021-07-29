package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	repository_base "github.com/jasonstanleyyoman/doremonangis_be/repository"
	repo_gorm "github.com/jasonstanleyyoman/doremonangis_be/repository/gorm"
	"github.com/jasonstanleyyoman/doremonangis_be/request_response"
	"github.com/jasonstanleyyoman/doremonangis_be/utils"
)

type StoreController struct {
	repository_base.MasterRepo
}

func NewStoreController() StoreController {
	repo := repo_gorm.GetRepo()
	return StoreController{
		repo,
	}
}

func (cont *StoreController) GetAllStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		stores := cont.GetStoreRepo().GetAllStore()
		c.JSON(http.StatusOK, stores)
	}
}

func (cont *StoreController) GetStoreInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		storeId, errConverting := utils.StringToUint(c.Param("id"))
		if errConverting != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Cannot convert " + c.Param("id") + "to uint",
			})
			return
		}
		store, errNotFound := cont.GetStoreRepo().GetStoreInfo(storeId)
		if errNotFound != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Store not found",
			})
			return
		}

		c.JSON(http.StatusOK, store)

	}
}

func (cont *StoreController) AddStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		var storeRequest request_response.AddStoreRequest

		if err := c.ShouldBindJSON(&storeRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		storeCreated, _ := cont.GetStoreRepo().AddStore(storeRequest.TransformToStore())
		storeInfo, _ := cont.GetStoreRepo().GetStoreInfo(storeCreated.ID)
		c.JSON(http.StatusOK, storeInfo)
	}
}

func (cont *StoreController) DeleteStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		storeId, errConverting := utils.StringToUint(c.Param("id"))
		if errConverting != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Cannot convert " + c.Param("id") + "to uint",
			})
			return
		}
		err := cont.GetStoreRepo().RemoveStore(storeId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully delete store",
		})

	}
}

func (cont *StoreController) AddStock() gin.HandlerFunc {
	return func(c *gin.Context) {
		var addStockRequest request_response.AlterStockRequest
		if err := c.ShouldBindJSON(&addStockRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if addStockRequest.Amount < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Cannot add stock with negative amount",
			})
			return
		}
		if err := cont.GetStoreRepo().AddStock(addStockRequest.StoreId, addStockRequest.DorayakiId, addStockRequest.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		store, _ := cont.GetStoreRepo().GetStoreInfo(addStockRequest.StoreId)
		c.JSON(http.StatusOK, store)
	}
}

func (cont *StoreController) RemoveStock() gin.HandlerFunc {
	return func(c *gin.Context) {
		var removeStockRequest request_response.AlterStockRequest
		if err := c.ShouldBindJSON(&removeStockRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if removeStockRequest.Amount < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Cannot remove stock with negative amount",
			})
			return
		}

		if err := cont.GetStoreRepo().RemoveStock(removeStockRequest.StoreId, removeStockRequest.DorayakiId, removeStockRequest.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		store, _ := cont.GetStoreRepo().GetStoreInfo(removeStockRequest.StoreId)
		c.JSON(http.StatusOK, store)
	}
}

func (cont *StoreController) MoveStock() gin.HandlerFunc {
	return func(c *gin.Context) {
		var moveStockRequest request_response.MoveStockRequest
		if err := c.ShouldBindJSON(&moveStockRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err := cont.GetStoreRepo().MoveStock(moveStockRequest.Source, moveStockRequest.Dest, moveStockRequest.DorayakiId, moveStockRequest.Amount)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		sourceStore, _ := cont.GetStoreRepo().GetStoreInfo(moveStockRequest.Source)
		destStore, _ := cont.GetStoreRepo().GetStoreInfo(moveStockRequest.Dest)

		c.JSON(http.StatusOK, request_response.MoveStockResponse{
			Source: sourceStore,
			Dest:   destStore,
		})

	}
}
