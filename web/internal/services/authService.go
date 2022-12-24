package services

import (
	"coursework/web/internal/models"
	"coursework/web/internal/repositories"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	salt      = "fdfsas12dfdsdv4"
	signInKey = "kjngjksdngn"
	tokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

type AuthorizationService struct {
	repo repositories.AuthRepo
}

func NewAuthService(repo repositories.AuthRepo) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) CreateUser(user models.User) (int, error) {
	user.Password = generatePassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthorizationService) GenerateToken(username string, password string) (string, error) {
	var user models.User
	var err error
	user, err = s.repo.GetUser(username, generatePassword(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})
	return token.SignedString([]byte(signInKey))
}

func (s *AuthorizationService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid method")
			}

			return []byte(signInKey), nil
		})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
