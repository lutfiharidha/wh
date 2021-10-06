package dto

import (
	"github.com/google/uuid"
)

type RackUpdateDTO struct {
	ID       uuid.UUID `json:"id" form:"id"`
	ZoneID   uuid.UUID `json:"zone_id" form:"product_name" binding:"required"`
	Aisle    string    `json:"aisle" form:"aisle"`
	Rack     string    `json:"rack" form:"rack"`
	Level    string    `json:"level" form:"level"`
	Position string    `json:"position" form:"position"`
}

type RackCreateDTO struct {
	ID       uuid.UUID `json:"id" form:"id"`
	ZoneID   uuid.UUID `json:"zone_id" form:"product_name" binding:"required"`
	Aisle    string    `json:"aisle" form:"aisle"`
	Rack     string    `json:"rack" form:"rack"`
	Level    string    `json:"level" form:"level"`
	Position string    `json:"position" form:"position"`
}

type RackRestoreDTO struct {
	ID uuid.UUID `json:"id" form:"id"`
}
