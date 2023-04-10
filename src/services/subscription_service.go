package services

import (
	"github.com/idprm/go-telenor-pk/src/domain/entity"
	"github.com/idprm/go-telenor-pk/src/domain/repository"
)

type SubscriptionService struct {
	subscriptionRepo repository.ISubscriptionRepository
}

func NewSubscriptionService(subscriptionRepo repository.ISubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepo: subscriptionRepo,
	}
}

type ISubscriptionService interface {
	GetActiveSubscription(int, string) bool
	GetSubscription(int, string) bool
	SaveSubscription(*entity.Subscription) error
	UpdateSubscription(*entity.Subscription) error
}

func (s *SubscriptionService) GetActiveSubscription(serviceId int, msisdn string) bool {
	count, _ := s.subscriptionRepo.CountActive(serviceId, msisdn)
	return count > 0
}

func (s *SubscriptionService) GetSubscription(serviceId int, msisdn string) bool {
	count, _ := s.subscriptionRepo.Count(serviceId, msisdn)
	return count > 0
}

func (s *SubscriptionService) SaveSubscription(sub *entity.Subscription) error {
	return nil
}

func (s *SubscriptionService) UpdateSubscription(sub *entity.Subscription) error {
	return nil
}
