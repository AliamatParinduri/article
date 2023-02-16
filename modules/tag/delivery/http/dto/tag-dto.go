package dto

type TagCreateDTO struct {
	Name string `json:"name" validate:"required,min=3,max=25"`
}
type TagUpdateDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required,min=3,max=25"`
}
