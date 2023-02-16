package usecase

import (
	"article_app/entity"
	"article_app/modules/tag/delivery/http/dto"
	repository "article_app/modules/tag/repository/postgres"
	"github.com/mashingan/smapping"
	"log"
)

type TagUsecase interface {
	GetTags() ([]entity.Tag, error)
	GetTagByID(id string) (*entity.Tag, error)
	CreateTag(tag *dto.TagDTO) (*entity.Tag, error)
	UpdateTag(id string, tag *dto.TagDTO) (*entity.Tag, error)
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

func (s *service) CreateTag(tag *dto.TagDTO) (*entity.Tag, error) {
	var u = entity.Tag{}
	err := smapping.FillStruct(&u, smapping.MapFields(&tag))
	if err != nil {
		log.Fatalf("failed map: %v", err)
	}
	return repo.CreateTag(&u)
}

func (s *service) UpdateTag(id string, tag *dto.TagDTO) (*entity.Tag, error) {
	var u = entity.Tag{}
	err := smapping.FillStruct(&u, smapping.MapFields(&tag))
	if err != nil {
		log.Fatalf("failed map: %v", err)
	}
	return repo.UpdateTag(id, &u)
}

func (s *service) DeleteTag(id string) error {
	return repo.DeleteTag(id)
}
