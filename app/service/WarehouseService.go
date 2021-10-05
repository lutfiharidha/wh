package service

import (
	"ginWeb/app/dto"
	"ginWeb/app/model"
	"ginWeb/app/repository"
	"log"

	"github.com/mashingan/smapping"
)

type WarehouseService interface {
	Insert(b dto.WarehouseCreateDTO) model.Warehouse
	Update(b dto.WarehouseUpdateDTO) model.Warehouse
	Restore(b dto.WarehouseRestoreDTO) model.Warehouse
	Delete(b model.Warehouse)
	DeletePermanent(b model.Warehouse)
	All() []model.Warehouse
	Deleted() []model.Warehouse
	FindByID(warehouseID string) model.Warehouse
	// IsAllowedToEdit(userID string, bookID uint64) bool
}

type warehouseService struct {
	warehouseRepository repository.WarehouseRepository
}

func NewWarehouseService(warehouseRepo repository.WarehouseRepository) WarehouseService {
	return &warehouseService{
		warehouseRepository: warehouseRepo,
	}
}

func (service *warehouseService) Insert(b dto.WarehouseCreateDTO) model.Warehouse {
	warehouse := model.Warehouse{}
	err := smapping.FillStruct(&warehouse, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.warehouseRepository.InsertWarehouse(warehouse)
	return res
}

func (service *warehouseService) Update(b dto.WarehouseUpdateDTO) model.Warehouse {
	warehouse := model.Warehouse{}
	err := smapping.FillStruct(&warehouse, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.warehouseRepository.UpdateWarehouse(warehouse)
	return res
}

func (service *warehouseService) Delete(b model.Warehouse) {
	service.warehouseRepository.DeleteWarehouse(b)
}

func (service *warehouseService) All() []model.Warehouse {
	return service.warehouseRepository.AllWarehouse()
}

func (service *warehouseService) FindByID(warehouseID string) model.Warehouse {
	return service.warehouseRepository.FindWarehouseByID(warehouseID)
}

func (service *warehouseService) Deleted() []model.Warehouse {
	return service.warehouseRepository.DeletedWarehouse()
}

func (service *warehouseService) Restore(b dto.WarehouseRestoreDTO) model.Warehouse {
	warehouse := model.Warehouse{}
	err := smapping.FillStruct(&warehouse, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.warehouseRepository.RestoreWarehouse(warehouse)
	return res
}

func (service *warehouseService) DeletePermanent(b model.Warehouse) {
	service.warehouseRepository.DeletePermanentWarehouse(b)
}

// func (service *warehouseService) IsAllowedToEdit(userID string, bookID uint64) bool {
// 	b := service.warehouseRepository.FindBookByID(warehouseID)
// 	id := fmt.Sprintf("%v", b.warehouseID)
// 	return userID == id
// }
