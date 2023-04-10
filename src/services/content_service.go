package services

import (
	"github.com/idprm/go-telenor-pk/src/domain/entity"
	"github.com/idprm/go-telenor-pk/src/domain/repository"
)

type ContentService struct {
	contentRepo repository.IContentRepository
}

type IContentService interface {
	GetContent(int, string) (*entity.Content, error)
}

func NewContentService(contentRepo repository.IContentRepository) *ContentService {
	return &ContentService{
		contentRepo: contentRepo,
	}
}

func (s *ContentService) GetContent(serviceId int, name string) (*entity.Content, error) {
	result, err := s.contentRepo.Get(serviceId, name)
	if err != nil {
		return nil, err
	}

	var content entity.Content

	if result != nil {
		content = entity.Content{
			ID:        result.ID,
			ServiceID: result.ServiceID,
			Name:      result.Name,
			Value:     result.Value,
		}
	}
	return &content, nil
}
