package service

import (
	"ginWeb/app/dto"
	"ginWeb/app/model"
	"ginWeb/app/repository"
	"log"

	"github.com/mashingan/smapping"
)

type RackService interface {
	Insert(rak dto.RackCreateDTO) model.Rack
	Update(rak dto.RackUpdateDTO) model.Rack
	Restore(rak dto.RackRestoreDTO) model.Rack
	Delete(rak model.Rack) model.Rack
	DeletePermanent(rak model.Rack)
	All() []model.Rack
	Deleted() []model.Rack
	FindByID(rackID string) model.Rack
	// IsAllowedToEdit(userID string, bookID uint64) bool
}

type rackService struct {
	rackRepository repository.RackRepository
}

func NewRackService(rackRepo repository.RackRepository) RackService {
	return &rackService{
		rackRepository: rackRepo,
	}
}

func (service *rackService) Insert(zon dto.RackCreateDTO) model.Rack {
	rack := model.Rack{}
	err := smapping.FillStruct(&rack, smapping.MapFields(&zon))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.rackRepository.InsertRack(rack)
	return res
}

func (service *rackService) Update(rak dto.RackUpdateDTO) model.Rack {
	rack := model.Rack{}
	err := smapping.FillStruct(&rack, smapping.MapFields(&rak))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.rackRepository.UpdateRack(rack)
	return res
}

func (service *rackService) Delete(rak model.Rack) model.Rack {
	return service.rackRepository.DeleteRack(rak)
}

func (service *rackService) All() []model.Rack {
	return service.rackRepository.AllRack()
}

func (service *rackService) FindByID(rackID string) model.Rack {
	return service.rackRepository.FindRackByID(rackID)
}

func (service *rackService) Deleted() []model.Rack {
	return service.rackRepository.DeletedRack()
}

func (service *rackService) Restore(rak dto.RackRestoreDTO) model.Rack {
	rack := model.Rack{}
	err := smapping.FillStruct(&rack, smapping.MapFields(&rak))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.rackRepository.RestoreRack(rack)
	return res
}

func (service *rackService) DeletePermanent(rak model.Rack) {
	service.rackRepository.DeletePermanentRack(rak)
}
