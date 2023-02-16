package usecase

import (
	"article_app/entity"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTUsecase interface {
	GenerateToken(user *entity.User) string
	Validate(token string) (*jwt.Token, error)
}

type JWTCustomClaims struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	IsAdmin  bool   `json:-`
	jwt.StandardClaims
}

type JwtUsecase struct {
	SecretKey string
	Issuer    string
}

func NewJWTUsecase() JWTUsecase {
	return &JwtUsecase{
		SecretKey: getSecretKey(),
		Issuer:    "AliamatParinduri-articleApp",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (j *JwtUsecase) GenerateToken(user *entity.User) string {
	claims := &JWTCustomClaims{
		user.ID,
		user.Name,
		user.Username,
		user.IsAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    j.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	// create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate encoded token pake secret key signing
	t, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		panic(err)
	}

	return t
}

func (j *JwtUsecase) Validate(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		//	return the secret signing key
		return []byte(j.SecretKey), nil
	})
}
