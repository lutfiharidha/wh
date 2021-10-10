package dto

type ZoneUpdateDTO struct {
	ID          string `json:"id" form:"id"`
	ZoneName    string `json:"zone_name" form:"zone_name" binding:"required"`
	WarehouseID string `json:"warehouse_id" form:"warehouse_id" binding:"required"`
}

type ZoneCreateDTO struct {
	ID          string `json:"id" form:"id"`
	ZoneName    string `json:"zone_name" form:"zone_name" binding:"required"`
	WarehouseID string `json:"warehouse_id" form:"warehouse_id" binding:"required"`
}

type ZoneRestoreDTO struct {
	ID string `json:"id" form:"id"`
}
