package usecase

import (
	"article_app/entity"
	repository "article_app/modules/auth/repository/postgres"
)

type AuthUsecase interface {
	LoginWithUsername(username string) (*entity.User, error)
	Register(user *entity.User) (*entity.User, error)
}

type service struct{}

var (
	repo repository.AuthRepository
)

func NewAuthUsecase(repp repository.AuthRepository) AuthUsecase {
	repo = repp
	return &service{}
}

func (s *service) Register(user *entity.User) (*entity.User, error) {
	return repo.Register(user)
}

func (s *service) LoginWithUsername(username string) (*entity.User, error) {
	return repo.LoginWithUsername(username)
}
