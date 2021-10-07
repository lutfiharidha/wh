package dto

import (
	"github.com/google/uuid"
)

type ProductLocationUpdateDTO struct {
	ID        uuid.UUID `json:"id" form:"id"`
	ProductID uuid.UUID `json:"product_id" form:"product_id" binding:"required" `
	RackID    uuid.UUID `json:"rack_id" form:"rack_id" binding:"required" `
	Stock     uint      `json:"stock" form:"stock" binding:"required"`
}

type ProductLocationCreateDTO struct {
	ID        uuid.UUID `json:"id" form:"id"`
	ProductID uuid.UUID `json:"product_id" form:"product_id" binding:"required" `
	RackID    uuid.UUID `json:"rack_id" form:"rack_id" binding:"required" `
	Stock     uint      `json:"stock" form:"stock" binding:"required"`
}

type ProductLocationRestoreDTO struct {
	ID uuid.UUID `json:"id" form:"id"`
}
