package usecase

import (
	"article_app/entity"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTUsecase interface {
	GenerateToken(user *entity.User) string
}

type JWTCustomClaims struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type jwtUsecase struct {
	secretKey string
	issuer    string
}

func NewJWTUsecase() JWTUsecase {
	return &jwtUsecase{
		secretKey: getSecretKey(),
		issuer:    "AliamatParinduri-articleApp",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (j *jwtUsecase) GenerateToken(user *entity.User) string {
	claims := &JWTCustomClaims{
		user.ID,
		user.Name,
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	// create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate encoded token pake secret key signing
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}

	return t
}
