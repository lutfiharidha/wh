package dto

import "github.com/google/uuid"

type WarehouseUpdateDTO struct {
	ID            string `json:"id" form:"id" gorm:""`
	WarehouseName string `json:"warehouse_name" form:"warehouse_name" binding:"required"`
	CompanyId     string `json:"company_id,omitempty" form:"company_id,omitempty"`
	Owner         string `json:"owner,omitempty" form:"owner,omitempty"`
	PhoneNumber   string `json:"phone_number" form:"phone_number" binding:"required"`
	Status        string `json:"status" form:"status" binding:"required"`
	Address       string `json:"address" form:"address" binding:"required"`
	City          string `json:"city" form:"city" binding:"required"`
	Country       string `json:"country" form:"country" binding:"required"`
	Geolocation   string `json:"geolocation" form:"geolocation" binding:"required"`
}

type WarehouseCreateDTO struct {
	ID            string `json:"id" form:"id"`
	WarehouseName string `json:"warehouse_name" form:"warehouse_name" binding:"required"`
	CompanyId     string `json:"company_id,omitempty" form:"company_id,omitempty"`
	Owner         string `json:"owner,omitempty" form:"owner,omitempty"`
	PhoneNumber   string `json:"phone_number" form:"phone_number" binding:"required"`
	Status        string `json:"status" form:"status" binding:"required"`
	Address       string `json:"address" form:"address" binding:"required"`
	City          string `json:"city" form:"city" binding:"required"`
	Country       string `json:"country" form:"country" binding:"required"`
	Geolocation   string `json:"geolocation" form:"geolocation" binding:"required"`
}

type WarehouseRestoreDTO struct {
	ID uuid.UUID `json:"id" form:"id"`
}
