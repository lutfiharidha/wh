package dto

type RackUpdateDTO struct {
	ID       string `json:"id" form:"id"`
	ZoneID   string `json:"zone_id" form:"zone_id" binding:"required"`
	Aisle    string `json:"aisle" form:"aisle"`
	Rack     string `json:"rack" form:"rack"`
	Level    string `json:"level" form:"level"`
	Position string `json:"position" form:"position"`
}

type RackCreateDTO struct {
	ID       string `json:"id" form:"id"`
	ZoneID   string `json:"zone_id" form:"zone_id" binding:"required"`
	Aisle    string `json:"aisle" form:"aisle"`
	Rack     string `json:"rack" form:"rack"`
	Level    string `json:"level" form:"level"`
	Position string `json:"position" form:"position"`
}

type RackRestoreDTO struct {
	ID string `json:"id" form:"id"`
}
