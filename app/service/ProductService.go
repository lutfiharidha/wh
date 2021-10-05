package service

import (
	"ginWeb/app/dto"
	"ginWeb/app/model"
	"ginWeb/app/repository"
	"log"

	"github.com/mashingan/smapping"
)

type ProductService interface {
	Insert(prod dto.ProductCreateDTO) model.Product
	Update(prod dto.ProductUpdateDTO) model.Product
	Restore(prod dto.ProductRestoreDTO) model.Product
	Delete(prod model.Product) model.Product
	DeletePermanent(prod model.Product)
	All() []model.Product
	Deleted() []model.Product
	FindByID(productID string) model.Product
	// IsAllowedToEdit(userID string, bookID uint64) bool
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepo,
	}
}

func (service *productService) Insert(prod dto.ProductCreateDTO) model.Product {
	product := model.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&prod))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.productRepository.InsertProduct(product)
	return res
}

func (service *productService) Update(prod dto.ProductUpdateDTO) model.Product {
	product := model.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&prod))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.productRepository.UpdateProduct(product)
	return res
}

func (service *productService) Delete(prod model.Product) model.Product {
	return service.productRepository.DeleteProduct(prod)
}

func (service *productService) All() []model.Product {
	return service.productRepository.AllProduct()
}

func (service *productService) FindByID(productID string) model.Product {
	return service.productRepository.FindProductByID(productID)
}

func (service *productService) Deleted() []model.Product {
	return service.productRepository.DeletedProduct()
}

func (service *productService) Restore(prod dto.ProductRestoreDTO) model.Product {
	product := model.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&prod))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.productRepository.RestoreProduct(product)
	return res
}

func (service *productService) DeletePermanent(prod model.Product) {
	service.productRepository.DeletePermanentProduct(prod)
}

// func (service *productService) IsAllowedToEdit(userID string, bookID uint64) bool {
// 	b := service.productRepository.FindBookByID(warehouseID)
// 	id := fmt.Sprintf("%v", b.warehouseID)
// 	return userID == id
// }
