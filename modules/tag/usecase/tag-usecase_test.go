package usecase

import (
	"article_app/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
	"time"
)

type mockRepository struct {
	mock.Mock
}

func TestGetTags(t *testing.T) {
	mockRepo := new(mockRepository)
	var number = 1

	tag := []entity.Tag{
		{
			ID:        number,
			Name:      "Go",
			Posts:     nil,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
	}
	mockRepo.On("GetTags").Return(tag, nil)

	testUsecase := NewTagUsecase(mockRepo)
	result, err := testUsecase.GetTags()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "Go", result[0].Name)
	assert.Nil(t, err)
}

func (mock *mockRepository) GetTags() ([]entity.Tag, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Tag), args.Error(1)
}

func TestFindByID(t *testing.T) {
	mockRepo := new(mockRepository)
	var number = "1"

	tag := entity.Tag{
		ID:        1,
		Name:      "Go",
		Posts:     nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
	mockRepo.On("GetTagByID").Return(&tag, nil)

	testUsecase := NewTagUsecase(mockRepo)
	result, err := testUsecase.GetTagByID(number)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "Go", result.Name)
	assert.Nil(t, err)
}

func (mock *mockRepository) GetTagByID(id string) (*entity.Tag, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Tag), args.Error(1)
}

func TestCreateTag(t *testing.T) {
	mockRepo := new(mockRepository)
	tag := entity.Tag{
		ID:        1,
		Name:      "Go",
		Posts:     nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}

	mockRepo.On("CreateTag").Return(&tag, nil)

	testUsecase := NewTagUsecase(mockRepo)
	result, err := testUsecase.CreateTag(&tag)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "Go", result.Name)
	assert.Nil(t, err)
}

func (mock *mockRepository) CreateTag(tag *entity.Tag) (*entity.Tag, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Tag), args.Error(1)
}

func TestUpdateTag(t *testing.T) {
	mockRepo := new(mockRepository)
	var number = "1"

	tag := entity.Tag{
		ID:        1,
		Name:      "Javascript",
		Posts:     nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}

	mockRepo.On("UpdateTag").Return(&tag, nil)

	testUsecase := NewTagUsecase(mockRepo)
	result, err := testUsecase.UpdateTag(number, &tag)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "Javascript", result.Name)
	assert.Nil(t, err)
}

func (mock *mockRepository) UpdateTag(id string, tag *entity.Tag) (*entity.Tag, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Tag), args.Error(1)
}

func TestDeleteTag(t *testing.T) {
	mockRepo := new(mockRepository)
	var number = "1"

	mockRepo.On("DeleteTag").Return(nil)

	testUsecase := NewTagUsecase(mockRepo)
	err := testUsecase.DeleteTag(number)

	mockRepo.AssertExpectations(t)

	assert.Nil(t, err)
}

func (mock *mockRepository) DeleteTag(id string) error {
	args := mock.Called()
	return args.Error(0)
}
