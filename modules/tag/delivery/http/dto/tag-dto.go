package dto

type TagDTO struct {
	Name string `json:"name" validate:"required,min=3,max=25"`
}
