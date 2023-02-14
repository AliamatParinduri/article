package postgres

import (
	"article_app/entity"
	"gorm.io/gorm"
)

type Server struct {
	Conn *gorm.DB
}

type AuthRepository interface {
	LoginWithUsername(user string) (*entity.User, error)
	Register(user *entity.User) (*entity.User, error)
}

func NewAuthRepository(conn *gorm.DB) AuthRepository {
	return &Server{conn}
}

func (s *Server) Register(user *entity.User) (*entity.User, error) {
	if err := s.Conn.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) LoginWithUsername(username string) (*entity.User, error) {
	var user = &entity.User{}
	if err := s.Conn.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
