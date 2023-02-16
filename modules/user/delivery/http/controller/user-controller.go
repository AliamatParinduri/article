package controller

import (
	"article_app/entity"
	"article_app/helper"
	"article_app/modules/user/delivery/http/dto"
	"article_app/modules/user/usecase"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserController interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userController struct{}

var (
	err error

	validate    *validator.Validate
	userUsecase usecase.UserUsecase
)

func NewUserController(useCase usecase.UserUsecase) UserController {
	userUsecase = useCase
	return &userController{}
}

func (c *userController) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []entity.User
	if users, err = userUsecase.GetUsers(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: err.Error()})
		return
	}

	if len(users) < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "User Not Found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success get all user",
		"data":    users,
	})
}

func (c *userController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user *entity.User
	params := mux.Vars(r)
	id := params["id"]

	if user, err = userUsecase.GetUserByID(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success get user",
		"data":    user,
	})
}

func (c *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user dto.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	defer r.Body.Close()

	message := helper.CustomeValidateError(user)
	if message != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: message})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Failed to hash password"})
		return
	}

	user.Password = string(hash)
	result, err := userUsecase.CreateUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error add user"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success add user",
		"data": map[string]interface{}{
			"id":       result.ID,
			"name":     result.Name,
			"username": result.Username,
		},
	})
}

func (c *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user dto.UserDTO
	params := mux.Vars(r)
	id := params["id"]

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	defer r.Body.Close()

	message := helper.CustomeValidateError(user)
	if message != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: message})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Failed to hash password"})
		return
	}

	user.Password = string(hash)
	result, err := userUsecase.UpdateUser(id, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error: " + err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success update user",
		"data": map[string]interface{}{
			"id":       result.ID,
			"name":     result.Name,
			"username": result.Username,
		},
	})
}

func (c *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	if err := userUsecase.DeleteUser(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error delete user"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success delete user",
	})
}
