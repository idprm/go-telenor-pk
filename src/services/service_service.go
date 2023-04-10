package services

import (
	"github.com/idprm/go-telenor-pk/src/domain/entity"
	"github.com/idprm/go-telenor-pk/src/domain/repository"
)

type ServiceService struct {
	serviceRepo repository.IServiceRepository
}

type IServiceService interface {
	GetService(string, string) bool
	GetServiceId(int) (*entity.Service, error)
	GetServiceByCode(string) (*entity.Service, error)
}

func NewServiceService(serviceRepo repository.IServiceRepository) *ServiceService {
	return &ServiceService{
		serviceRepo: serviceRepo,
	}
}

func (s *ServiceService) GetService(code string, pkg string) bool {
	count, _ := s.serviceRepo.Count(code, pkg)
	return count > 0
}

func (s *ServiceService) GetServiceId(id int) (*entity.Service, error) {
	result, err := s.serviceRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	var srv entity.Service

	if result != nil {
		srv = entity.Service{
			ID:              result.ID,
			Category:        result.Category,
			Code:            result.Code,
			Name:            result.Name,
			Package:         result.Package,
			AuthUser:        result.AuthUser,
			AuthPass:        result.AuthPass,
			PartnerID:       result.PartnerID,
			ProductID:       result.ProductID,
			TransID:         result.TransID,
			CorrelationID:   result.CorrelationID,
			Price:           result.Price,
			RenewalDay:      result.RenewalDay,
			TrialDay:        result.TrialDay,
			UrlHost:         result.UrlHost,
			UrlCallback:     result.UrlCallback,
			UrlNotifSub:     result.UrlNotifSub,
			UrlNotifUnsub:   result.UrlNotifUnsub,
			UrlNotifRenewal: result.UrlNotifRenewal,
			UrlPostback:     result.UrlPostback,
		}
	}
	return &srv, nil
}

func (s *ServiceService) GetServiceByCode(code string) (*entity.Service, error) {
	result, err := s.serviceRepo.GetByCode(code)
	if err != nil {
		return nil, err
	}

	var srv entity.Service

	if result != nil {
		srv = entity.Service{
			ID:              result.ID,
			Category:        result.Category,
			Code:            result.Code,
			Name:            result.Name,
			Package:         result.Package,
			AuthUser:        result.AuthUser,
			AuthPass:        result.AuthPass,
			PartnerID:       result.PartnerID,
			ProductID:       result.ProductID,
			TransID:         result.TransID,
			CorrelationID:   result.CorrelationID,
			Price:           result.Price,
			RenewalDay:      result.RenewalDay,
			TrialDay:        result.TrialDay,
			UrlHost:         result.UrlHost,
			UrlCallback:     result.UrlCallback,
			UrlNotifSub:     result.UrlNotifSub,
			UrlNotifUnsub:   result.UrlNotifUnsub,
			UrlNotifRenewal: result.UrlNotifRenewal,
			UrlPostback:     result.UrlPostback,
		}
	}
	return &srv, nil
}
