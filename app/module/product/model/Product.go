package model

import (
	"time"

	models "ginWeb/app/module/warehouse/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	// gorm.Model
	ID              uuid.UUID      `gorm:"uniqueIndex" json:"id"`
	ProductName     string         `json:"product_name" gorm:"not null"`
	WarehouseID     uuid.UUID      `gorm:"index;not null" json:"warehouse_id"`
	MasterProductID *uuid.UUID     `json:"master_product_id"`
	SKU             string         `json:"sku" gorm:"not null"`
	Description     string         `json:"description"`
	Price           float32        `json:"price" gorm:"not null"`
	Image           string         `json:"image" gorm:"not null"`
	Status          uint8          `json:"status" gorm:"not null"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index"  json:"deleted_at"`

	Warehouse *models.Warehouse `gorm:"references:ID;foreignKey:WarehouseID"`
}
