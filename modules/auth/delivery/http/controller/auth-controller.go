package controller

import (
	"article_app/entity"
	"article_app/helper"
	"article_app/modules/Auth/usecase"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	jwtUsecase usecase.JWTUsecase
}

var (
	authUsecase usecase.AuthUsecase
)

func NewAuthController(useCase usecase.AuthUsecase, jwtUsecase usecase.JWTUsecase) AuthController {
	authUsecase = useCase
	return &authController{
		jwtUsecase: jwtUsecase,
	}
}

func (c *authController) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	defer r.Body.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Failed to hash password"})
		return
	}

	user.Password = string(hash)
	result, err := authUsecase.Register(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error register user"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success register user",
		"data": map[string]interface{}{
			"id":       result.ID,
			"name":     result.Name,
			"username": result.Username,
		},
	})
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userInput entity.User

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	defer r.Body.Close()

	user, err := authUsecase.LoginWithUsername(userInput.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Invalid username or password"})
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Invalid username or password"})
		return
	}

	// generate token
	token := c.jwtUsecase.GenerateToken(user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success login",
		"token":   token,
		"data": map[string]interface{}{
			"id":       user.ID,
			"name":     user.Name,
			"username": user.Username,
		},
	})
}
