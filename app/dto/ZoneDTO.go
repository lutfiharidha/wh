package dto

import (
	"github.com/google/uuid"
)

type ZoneUpdateDTO struct {
	ID          uuid.UUID `json:"id" form:"id"`
	ZoneName    string    `json:"zone_name" form:"zone_name" binding:"required"`
	WarehouseID uuid.UUID `json:"warehouse_id" form:"warehouse_id" binding:"required"`
}

type ZoneCreateDTO struct {
	ID          uuid.UUID `json:"id" form:"id"`
	ZoneName    string    `json:"zone_name" form:"zone_name" binding:"required"`
	WarehouseID uuid.UUID `json:"warehouse_id" form:"warehouse_id" binding:"required"`
}

type ZoneRestoreDTO struct {
	ID uuid.UUID `json:"id" form:"id"`
}
