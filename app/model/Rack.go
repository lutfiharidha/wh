package model

import (
	// "ginWeb/app/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rack struct {
	ID        uuid.UUID      `gorm:"uniqueIndex" json:"id"`
	ZoneID    uuid.UUID      `gorm:"index" json:"zone_id"`
	Aisle     string         `json:"aisle"`
	Rack      string         `json:"rack"`
	Level     string         `json:"level"`
	Position  string         `json:"position"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"  json:"deleted_at"`

	Zone *Zone `gorm:"foreignKey:ZoneID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"zone"`
}

func (rack *Rack) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return
}
