package model

import (
	// "ginWeb/app/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductLocation struct {
	ID        uuid.UUID      `gorm:"uniqueIndex" json:"id"`
	ProductID uuid.UUID      `gorm:"index;not null" json:"-"`
	RackID    uuid.UUID      `gorm:"index;not null" json:"-" `
	Stock     uint           `json:"stock" gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"  json:"deleted_at"`

	Product Product `gorm:"foreignKey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"product"`
	Rack    Rack    `gorm:"foreignKey:RackID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"rack"`
}

func (productLocation *ProductLocation) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return
}
