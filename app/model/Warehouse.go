package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Warehouse struct {
	ID            uuid.UUID      `gorm:"uniqueIndex" json:"id"`
	WarehouseName string         `json:"warehouse_name" gorm:"not null"`
	CompanyId     uuid.UUID      `json:"company_id"`
	Owner         string         `json:"owner" gorm:"not null"`
	PhoneNumber   string         `json:"phone_number" gorm:"not null"`
	Status        uint8          `json:"status" gorm:"not null"`
	Address       string         `json:"address" gorm:"not null"`
	City          string         `json:"city" gorm:"not null"`
	Country       string         `json:"country" gorm:"not null"`
	Geolocation   string         `json:"geolocation" gorm:"not null"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"  json:"deleted_at"`
}

func (warehouse *Warehouse) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return
}
