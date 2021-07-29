package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	repository_base "github.com/jasonstanleyyoman/doremonangis_be/repository"
	repo_gorm "github.com/jasonstanleyyoman/doremonangis_be/repository/gorm"
	"github.com/jasonstanleyyoman/doremonangis_be/request_response"
	"github.com/jasonstanleyyoman/doremonangis_be/utils"
)

type DorayakiController struct {
	repository_base.MasterRepo
}

func NewDorayakiController() DorayakiController {
	repo := repo_gorm.GetRepo()
	return DorayakiController{
		repo,
	}
}

func (cont *DorayakiController) GetAllDorayaki() gin.HandlerFunc {
	return func(c *gin.Context) {
		allDorayaki := cont.GetDorayakiRepo().GetAllDorayaki()

		c.JSON(http.StatusOK, allDorayaki)
	}
}

func (cont *DorayakiController) GetDorayakiInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		dorayakiId, errConverting := utils.StringToUint(c.Param("id"))
		if errConverting != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Cannot convert " + c.Param("id") + " to uint",
			})
			return
		}
		dorayaki, errNotFound := cont.GetDorayakiRepo().GetDorayakiInfo(dorayakiId)
		if errNotFound != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Dorayaki not found",
			})
			return
		}

		c.JSON(http.StatusOK, dorayaki)

	}
}

func (cont *DorayakiController) AddDorayaki() gin.HandlerFunc {
	return func(c *gin.Context) {
		var addDorayakiRequest request_response.AddDorayakiRequest
		if err := c.ShouldBindJSON(&addDorayakiRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		dorayakiCreated, _ := cont.GetDorayakiRepo().AddDorayaki(addDorayakiRequest.TransformToDorayaki())
		dorayakiCreated, _ = cont.GetDorayakiRepo().GetDorayakiInfo(dorayakiCreated.ID)

		c.JSON(http.StatusOK, dorayakiCreated)
	}
}

func (cont *DorayakiController) DeleteDorayaki() gin.HandlerFunc {
	return func(c *gin.Context) {
		dorayakiId, errConverting := utils.StringToUint(c.Param("id"))
		if errConverting != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Cannot convert " + c.Param("id") + "to uint",
			})
			return
		}
		err := cont.GetDorayakiRepo().RemoveDorayaki(dorayakiId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully delete dorayaki",
		})
	}
}
