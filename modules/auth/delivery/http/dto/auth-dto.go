package dto

type LoginDTO struct {
	Username string `json:"username" validate:"required,min=3,max=15,alphanum"`
	Password string `json:"password,omitempty" validate:"required,min=4"`
}

type RegisterDTO struct {
	Name     string `json:"name" validate:"required,min=3"`
	Username string `json:"username" validate:"required,min=3,max=15,alphanum"`
	Password string `json:"password,omitempty" validate:"required,min=4"`
}
