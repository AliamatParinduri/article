package postgres

import (
	"article_app/entity"
	"errors"
	"gorm.io/gorm"
)

type Server struct {
	Conn *gorm.DB
}

type UserRepository interface {
	GetUsers() ([]entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(id string, user *entity.User) (*entity.User, error)
	DeleteUser(id string) error
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return &Server{conn}
}

func (s *Server) GetUsers() ([]entity.User, error) {
	var users []entity.User
	if err := s.Conn.Where("is_admin = false").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Server) GetUserByID(id string) (*entity.User, error) {
	var user *entity.User
	if err := s.Conn.Where("is_admin = false").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) CreateUser(user *entity.User) (*entity.User, error) {
	if err := s.Conn.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) UpdateUser(id string, user *entity.User) (*entity.User, error) {
	if rows := s.Conn.Where("is_admin = false").Where("is_admin = false").Where("id = ?", id).Updates(&user).RowsAffected; rows == 0 {
		return nil, errors.New("Failed update data user, record not found")
	}
	return user, nil
}

func (s *Server) DeleteUser(id string) error {
	var user entity.User
	if rows := s.Conn.Where("id = ?", id).Delete(&user).RowsAffected; rows == 0 {
		return errors.New("Failed delete data user")
	}

	return nil
}
