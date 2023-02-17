package usecase

import (
	"article_app/entity"
	repository "article_app/modules/tag/repository/postgres"
)

type TagUsecase interface {
	GetTags() ([]entity.Tag, error)
	GetTagByID(id string) (*entity.Tag, error)
	CreateTag(tag *entity.Tag) (*entity.Tag, error)
	UpdateTag(id string, tag *entity.Tag) (*entity.Tag, error)
	DeleteTag(id string) error
}

type service struct{}

var (
	repo repository.TagRepository
)

func NewTagUsecase(repp repository.TagRepository) TagUsecase {
	repo = repp
	return &service{}
}

func (s *service) GetTags() ([]entity.Tag, error) {
	return repo.GetTags()
}

func (s *service) GetTagByID(id string) (*entity.Tag, error) {
	return repo.GetTagByID(id)
}

func (s *service) CreateTag(tag *entity.Tag) (*entity.Tag, error) {
	return repo.CreateTag(tag)
}

func (s *service) UpdateTag(id string, tag *entity.Tag) (*entity.Tag, error) {
	return repo.UpdateTag(id, tag)
}

func (s *service) DeleteTag(id string) error {
	return repo.DeleteTag(id)
}
