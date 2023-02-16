package postgres

import (
	"article_app/entity"
	"gorm.io/gorm"
)

type Server struct {
	Conn *gorm.DB
}

type PostRepository interface {
	GetPosts() ([]entity.Post, error)
	GetPostByUser(id string) ([]entity.Post, error)
	GetPostByID(id string) (*entity.Post, error)
	CreatePost(post *entity.Post) (*entity.Post, error)
	UpdatePost(id int, Post *entity.Post) (*entity.Post, error)
	DeletePost(id int) error
}

func NewPostRepository(conn *gorm.DB) PostRepository {
	return &Server{conn}
}

func (s *Server) GetPosts() ([]entity.Post, error) {
	var posts []entity.Post
	if err := s.Conn.Preload("User").Preload("Tags").Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *Server) GetPostByID(id string) (*entity.Post, error) {
	var post *entity.Post
	if err := s.Conn.Where("id = ?", id).Preload("User").Preload("Tags").First(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}
func (s *Server) GetPostByUser(id string) ([]entity.Post, error) {
	var post []entity.Post
	if err := s.Conn.Where("user_id = ?", id).Preload("User").Preload("Tags").First(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (s *Server) CreatePost(post *entity.Post) (*entity.Post, error) {
	if err := s.Conn.Create(&post).Association("Tags").Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (s *Server) UpdatePost(id int, post *entity.Post) (*entity.Post, error) {
	var newPost = *post

	s.DeletePost(id)

	s.Conn.Create(&newPost)

	return &newPost, nil
}

func (s *Server) DeletePost(id int) error {
	var post entity.Post
	//// find tags related to post
	s.Conn.Where("id = ?", id).Preload("Tags").First(&post)

	////// delete tags related to post
	if err := s.Conn.Model(&entity.Post{ID: id}).Association("Tags").Clear(); err != nil {
		return err
	}

	s.Conn.Where("id = ?", id).Unscoped().Delete(&post)
	return nil
}
