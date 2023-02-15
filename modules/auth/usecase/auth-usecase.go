package usecase

import (
	"article_app/entity"
	"article_app/modules/auth/delivery/http/dto"
	repository "article_app/modules/auth/repository/postgres"
	"github.com/mashingan/smapping"
	"log"
)

type AuthUsecase interface {
	LoginWithUsername(username string) (*entity.User, error)
	Register(user *dto.AuthDTO) (*entity.User, error)
}

type service struct{}

var (
	repo repository.AuthRepository
)

func NewAuthUsecase(repp repository.AuthRepository) AuthUsecase {
	repo = repp
	return &service{}
}

func (s *service) Register(user *dto.AuthDTO) (*entity.User, error) {
	var u = entity.User{}
	err := smapping.FillStruct(&u, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed map: %v", err)
	}
	return repo.Register(&u)
}

func (s *service) LoginWithUsername(username string) (*entity.User, error) {
	return repo.LoginWithUsername(username)
}
