package usecase

import (
	"article_app/entity"
	"article_app/modules/post/delivery/http/dto"
	repository "article_app/modules/post/repository/postgres"
	"log"

	"github.com/mashingan/smapping"
)

type PostUsecase interface {
	GetPosts() ([]entity.Post, error)
	GetPostByUser(id string) ([]entity.Post, error)
	GetPostByID(id string) (*entity.Post, error)
	CreatePost(post *dto.PostDTO) (*entity.Post, error)
	UpdatePost(id int, post *dto.PostDTO) (*entity.Post, error)
	DeletePost(id int) error
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostUsecase(repp repository.PostRepository) PostUsecase {
	repo = repp
	return &service{}
}

func (s *service) GetPosts() ([]entity.Post, error) {
	return repo.GetPosts()
}

func (s *service) GetPostByID(id string) (*entity.Post, error) {
	return repo.GetPostByID(id)
}
func (s *service) GetPostByUser(id string) ([]entity.Post, error) {
	return repo.GetPostByUser(id)
}

func (s *service) CreatePost(post *dto.PostDTO) (*entity.Post, error) {
	var u = entity.Post{}
	err := smapping.FillStruct(&u, smapping.MapFields(&post))
	if err != nil {
		log.Fatalf("failed map: %v", err)
	}
	return repo.CreatePost(&u)
}

func (s *service) UpdatePost(id int, post *dto.PostDTO) (*entity.Post, error) {
	var u = entity.Post{}
	err := smapping.FillStruct(&u, smapping.MapFields(&post))
	if err != nil {
		log.Fatalf("failed map: %v", err)
	}
	return repo.UpdatePost(id, &u)
}

func (s *service) DeletePost(id int) error {
	return repo.DeletePost(id)
}
