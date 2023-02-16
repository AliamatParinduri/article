package postgres

import (
	"article_app/entity"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type Server struct {
	Conn *gorm.DB
}

type TagRepository interface {
	GetTags() ([]entity.Tag, error)
	GetTagByID(id string) (*entity.Tag, error)
	CreateTag(tag *entity.Tag) (*entity.Tag, error)
	UpdateTag(id string, tag *entity.Tag) (*entity.Tag, error)
	DeleteTag(id string) error
}

func NewTagRepository(conn *gorm.DB) TagRepository {
	return &Server{conn}
}

func (s *Server) GetTags() ([]entity.Tag, error) {
	var tags []entity.Tag
	if err := s.Conn.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (s *Server) GetTagByID(id string) (*entity.Tag, error) {
	var tag *entity.Tag
	if err := s.Conn.Where("id = ?", id).First(&tag).Error; err != nil {
		return nil, err
	}

	return tag, nil
}

func (s *Server) CreateTag(tag *entity.Tag) (*entity.Tag, error) {
	if err := s.Conn.Create(&tag).Error; err != nil {
		return nil, err
	}

	return tag, nil
}

func (s *Server) UpdateTag(id string, tag *entity.Tag) (*entity.Tag, error) {
	if rows := s.Conn.Where("id = ?", id).Updates(&tag).RowsAffected; rows == 0 {
		return nil, errors.New("Failed update data tag, record not found")
	}
	tag.ID, _ = strconv.Atoi(id)
	return tag, nil
}

func (s *Server) DeleteTag(id string) error {
	var tag entity.Tag
	if rows := s.Conn.Where("id = ?", id).Delete(&tag).RowsAffected; rows == 0 {
		return errors.New("Failed delete data tag")
	}

	return nil
}
