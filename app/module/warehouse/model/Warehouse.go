package models

import "ginWeb/app/config"

type Warehouse struct {
	config.Model

	WarehouseName string `json:"warehouse_name" gorm:"not null"`
	CompanyId     string `json:"company_id"`
	Owner         string `json:"owner" gorm:"not null"`
	PhoneNumber   string `json:"phone_number" gorm:"not null"`
	Status        uint8  `json:"status" gorm:"not null"`
	Address       string `json:"address" gorm:"not null"`
	City          string `json:"city" gorm:"not null"`
	Country       string `json:"country" gorm:"not null"`
	Geolocation   string `json:"geolocation" gorm:"not null"`
}
