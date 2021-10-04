package handler

import (
	models "ginWeb/app/module/warehouse/model"

	"gorm.io/gorm"
)

func GetWarehouses(db *gorm.DB) ([]models.Warehouse, error) {
	warehouses := []models.Warehouse{}
	query := db.Find(&warehouses)
	if err := query.Error; err != nil {
		return warehouses, err
	}

	return warehouses, nil
}
