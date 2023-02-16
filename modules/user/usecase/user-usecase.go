package usecase

import (
	"article_app/entity"
	"article_app/modules/user/delivery/http/dto"
	repository "article_app/modules/user/repository/postgres"
	"github.com/mashingan/smapping"
	"log"
)

type UserUsecase interface {
	GetUsers() ([]entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	CreateUser(user *dto.UserDTO) (*entity.User, error)
	UpdateUser(id string, user *dto.UserDTO) (*entity.User, error)
	DeleteUser(id string) error
}

type service struct{}

var (
	repo repository.UserRepository
)

func NewUserUsecase(repp repository.UserRepository) UserUsecase {
	repo = repp
	return &service{}
}

func (s *service) GetUsers() ([]entity.User, error) {
	return repo.GetUsers()
}

func (s *service) GetUserByID(id string) (*entity.User, error) {
	return repo.GetUserByID(id)
}

func (s *service) CreateUser(user *dto.UserDTO) (*entity.User, error) {
	var u = entity.User{}
	err := smapping.FillStruct(&u, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed map: %v", err)
	}
	return repo.CreateUser(&u)
}

func (s *service) UpdateUser(id string, user *dto.UserDTO) (*entity.User, error) {
	var u = entity.User{}
	err := smapping.FillStruct(&u, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed map: %v", err)
	}
	return repo.UpdateUser(id, &u)
}

func (s *service) DeleteUser(id string) error {
	return repo.DeleteUser(id)
}
