package service

import (
	"ginWeb/app/dto"
	"ginWeb/app/model"
	"ginWeb/app/repository"
	"log"

	"github.com/mashingan/smapping"
)

type ProductLocationService interface {
	Insert(prod dto.ProductLocationCreateDTO) model.ProductLocation
	Update(prod dto.ProductLocationUpdateDTO) model.ProductLocation
	Restore(prod dto.ProductLocationRestoreDTO) model.ProductLocation
	Delete(prod model.ProductLocation) model.ProductLocation
	DeletePermanent(prod model.ProductLocation)
	All() []model.ProductLocation
	Deleted() []model.ProductLocation
	FindByID(productID string) model.ProductLocation
	// IsAllowedToEdit(userID string, bookID uint64) bool
}

type productLocationService struct {
	productLocationRepository repository.ProductLocationRepository
}

func NewProductLocationService(productLocationRepo repository.ProductLocationRepository) ProductLocationService {
	return &productLocationService{
		productLocationRepository: productLocationRepo,
	}
}

func (service *productLocationService) Insert(prod dto.ProductLocationCreateDTO) model.ProductLocation {
	product := model.ProductLocation{}
	err := smapping.FillStruct(&product, smapping.MapFields(&prod))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.productLocationRepository.InsertProductLocation(product)
	return res
}

func (service *productLocationService) Update(prod dto.ProductLocationUpdateDTO) model.ProductLocation {
	product := model.ProductLocation{}
	err := smapping.FillStruct(&product, smapping.MapFields(&prod))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.productLocationRepository.UpdateProductLocation(product)
	return res
}

func (service *productLocationService) Delete(pl model.ProductLocation) model.ProductLocation {
	return service.productLocationRepository.DeleteProductLocation(pl)
}

func (service *productLocationService) All() []model.ProductLocation {
	return service.productLocationRepository.AllProductLocation()
}

func (service *productLocationService) FindByID(productLocationID string) model.ProductLocation {
	return service.productLocationRepository.FindProductLocationByID(productLocationID)
}

func (service *productLocationService) Deleted() []model.ProductLocation {
	return service.productLocationRepository.DeletedProductLocation()
}

func (service *productLocationService) Restore(pl dto.ProductLocationRestoreDTO) model.ProductLocation {
	product := model.ProductLocation{}
	err := smapping.FillStruct(&product, smapping.MapFields(&pl))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.productLocationRepository.RestoreProductLocation(product)
	return res
}

func (service *productLocationService) DeletePermanent(pl model.ProductLocation) {
	service.productLocationRepository.DeletePermanentProductLocation(pl)
}

// func (service *productLocationService) IsAllowedToEdit(userID string, bookID uint64) bool {
// 	b := service.productLocationRepository.FindBookByID(warehouseID)
// 	id := fmt.Sprintf("%v", b.warehouseID)
// 	return userID == id
// }
