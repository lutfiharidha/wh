package dto

import (
	"github.com/google/uuid"
)

type ProductUpdateDTO struct {
	ID              uuid.UUID  `json:"id" form:"id"`
	ProductName     string     `json:"product_name" form:"product_name" binding:"required"`
	WarehouseID     uuid.UUID  `json:"warehouse_id" form:"warehouse_id" binding:"required"`
	MasterProductID *uuid.UUID `json:"master_product_id" form:"master_product_id"`
	SKU             string     `json:"sku" form:"sku" binding:"required"`
	Description     string     `json:"description" form:"description"`
	Price           float32    `json:"price" form:"price" binding:"required"`
	Image           string     `json:"image" form:"image" binding:"required"`
	Status          uint8      `json:"status" form:"status" binding:"required"`
}

type ProductCreateDTO struct {
	ID              uuid.UUID  `json:"id" form:"id"`
	ProductName     string     `json:"product_name" form:"product_name" binding:"required"`
	WarehouseID     uuid.UUID  `json:"warehouse_id" form:"warehouse_id" binding:"required"`
	MasterProductID *uuid.UUID `json:"master_product_id" form:"master_product_id"`
	SKU             string     `json:"sku" form:"sku" binding:"required"`
	Description     string     `json:"description" form:"description"`
	Price           float32    `json:"price" form:"price" binding:"required"`
	Image           string     `json:"image" form:"image" binding:"required"`
	Status          uint8      `json:"status" form:"status" binding:"required"`
}

type ProductRestoreDTO struct {
	ID uuid.UUID `json:"id" form:"id"`
}
