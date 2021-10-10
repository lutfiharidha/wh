package dto

type ProductLocationUpdateDTO struct {
	ID        string `json:"id" form:"id"`
	ProductID string `json:"product_id" form:"product_id" binding:"required" `
	RackID    string `json:"rack_id" form:"rack_id" binding:"required" `
	Stock     uint   `json:"stock" form:"stock" binding:"required"`
}

type ProductLocationCreateDTO struct {
	ID        string `json:"id" form:"id"`
	ProductID string `json:"product_id" form:"product_id" binding:"required" `
	RackID    string `json:"rack_id" form:"rack_id" binding:"required" `
	Stock     uint   `json:"stock" form:"stock" binding:"required"`
}

type ProductLocationRestoreDTO struct {
	ID string `json:"id" form:"id"`
}
