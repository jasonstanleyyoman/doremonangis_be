package entity

import (
	"gorm.io/gorm"
)

type StoreDorayaki struct {
	StoreId    uint `gorm:"constraint:OnDelete:CASCADE,OnDelete:CASCADE;primaryKey;"`
	DorayakiId uint `gorm:"constraint:OnDelete:CASCADE,OnDelete:CASCADE;primaryKey;"`
	Amount     int
}

func (StoreDorayaki) TableName() string {
	return "store_dorayaki"
}
func (sd *StoreDorayaki) AfterSave(tx *gorm.DB) (err error) {
	if sd.Amount <= 0 {
		tx.Where("store_id = ? AND dorayaki_id = ?", sd.StoreId, sd.DorayakiId).Delete(&StoreDorayaki{})
	}
	return nil
}

func (sd *StoreDorayaki) AfterUpdate(tx *gorm.DB) (err error) {
	if sd.Amount <= 0 {
		tx.Where("store_id = ? AND dorayaki_id = ?", sd.StoreId, sd.DorayakiId).Delete(&StoreDorayaki{})
	}
	return nil
}
