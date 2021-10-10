package model

import (
	// "ginWeb/app/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID              string         `gorm:"uniqueIndex" json:"id"`
	ProductName     string         `json:"product_name" gorm:"not null"`
	WarehouseID     string         `gorm:"index;not null" json:"-"`
	MasterProductID *string        `json:"master_product_id"`
	SKU             string         `json:"sku" gorm:"not null"`
	Description     string         `json:"description"`
	Price           float32        `json:"price" gorm:"not null"`
	Image           string         `json:"image" gorm:"not null"`
	Status          uint8          `json:"status" gorm:"not null"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index"  json:"deleted_at"`

	Warehouse *Warehouse `gorm:"foreignKey:WarehouseID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"warehouse"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return
}
