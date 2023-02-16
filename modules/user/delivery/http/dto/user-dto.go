package dto

type UserDTO struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Username string `json:"username" validate:"required,min=3,max=25,alphanum"`
	Password string `json:"password,omitempty" validate:"required,min=4"`
}
