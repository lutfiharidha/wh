package service

import (
	"ginWeb/app/dto"
	"ginWeb/app/model"
	"ginWeb/app/repository"
	"log"

	"github.com/mashingan/smapping"
)

type ZoneService interface {
	Insert(zon dto.ZoneCreateDTO) model.Zone
	Update(zon dto.ZoneUpdateDTO) model.Zone
	Restore(zon dto.ZoneRestoreDTO) model.Zone
	Delete(zon model.Zone) model.Zone
	DeletePermanent(zon model.Zone)
	All() []model.Zone
	Deleted() []model.Zone
	FindByID(zoneID string) model.Zone
	// IsAllowedToEdit(userID string, bookID uint64) bool
}

type zoneService struct {
	zoneRepository repository.ZoneRepository
}

func NewZoneService(zoneRepo repository.ZoneRepository) ZoneService {
	return &zoneService{
		zoneRepository: zoneRepo,
	}
}

func (service *zoneService) Insert(zon dto.ZoneCreateDTO) model.Zone {
	zone := model.Zone{}
	err := smapping.FillStruct(&zone, smapping.MapFields(&zon))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.zoneRepository.InsertZone(zone)
	return res
}

func (service *zoneService) Update(zon dto.ZoneUpdateDTO) model.Zone {
	zone := model.Zone{}
	err := smapping.FillStruct(&zone, smapping.MapFields(&zon))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.zoneRepository.UpdateZone(zone)
	return res
}

func (service *zoneService) Delete(zon model.Zone) model.Zone {
	return service.zoneRepository.DeleteZone(zon)
}

func (service *zoneService) All() []model.Zone {
	return service.zoneRepository.AllZone()
}

func (service *zoneService) FindByID(zoneID string) model.Zone {
	return service.zoneRepository.FindZoneByID(zoneID)
}

func (service *zoneService) Deleted() []model.Zone {
	return service.zoneRepository.DeletedZone()
}

func (service *zoneService) Restore(zon dto.ZoneRestoreDTO) model.Zone {
	zone := model.Zone{}
	err := smapping.FillStruct(&zone, smapping.MapFields(&zon))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.zoneRepository.RestoreZone(zone)
	return res
}

func (service *zoneService) DeletePermanent(zon model.Zone) {
	service.zoneRepository.DeletePermanentZone(zon)
}
