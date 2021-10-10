package model

import (
	// "ginWeb/app/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Zone struct {
	ID          string         `gorm:"uniqueIndex" json:"id"`
	ZoneName    string         `json:"zone_name" gorm:"not null"`
	WarehouseID string         `json:"-" gorm:"index; not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"  json:"deleted_at"`

	Warehouse Warehouse `gorm:"foreignKey:WarehouseID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"warehouse"`
}

func (zone *Zone) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return
}
