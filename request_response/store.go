package request_response

import "github.com/jasonstanleyyoman/doremonangis_be/entity"

type AddStoreRequest struct {
	Name      string `json:"name" binding:"required,max=128"`
	Address   string `json:"address" binding:"required,max=256"`
	Kecamatan string `json:"kecamatan" binding:"required,max=128"`
	Province  string `json:"province" binding:"required,max=128"`
}

type AlterStockRequest struct {
	StoreId    uint `json:"store_id" binding:"required"`
	DorayakiId uint `json:"dorayaki_id" binding:"required"`
	Amount     int  `json:"amount" binding:"required,gt=0"`
}

type MoveStockRequest struct {
	Source     uint `json:"src" binding:"required"`
	Dest       uint `json:"dest" binding:"required"`
	DorayakiId uint `json:"dorayaki_id" binding:"required"`
	Amount     int  `json:"amount" binding:"required"`
}

type MoveStockResponse struct {
	Source entity.StoreWithDorayakiAmount `json:"src"`
	Dest   entity.StoreWithDorayakiAmount `json:"dest"`
}

func (req AddStoreRequest) TransformToStore() entity.Store {
	return entity.Store{
		Name:      req.Name,
		Address:   req.Address,
		Kecamatan: req.Kecamatan,
		Province:  req.Province,
	}
}
