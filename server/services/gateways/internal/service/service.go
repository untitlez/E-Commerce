package service

import (
	"errors"
	"server/services/gateways/internal/domain"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo domain.AuthRepository
}

type JWTClaims struct {
	jwt.RegisteredClaims
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func NewService(r domain.AuthRepository) *service {
	return &service{repo: r}
}

func (s *service) SignUp(body *domain.Auth) error {
	req := body.Username == "" || body.Password == ""

	if req {
		return errors.New("username or password invalid")
	}

	if _, err := s.repo.FindByUsername(body); err != nil {
		return errors.New("username already exists")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return err
	}

	newBody := &domain.Auth{
		Username: body.Username,
		Password: string(hashPassword),
	}

	if err := s.repo.Create(newBody); err != nil {
		return errors.New("fail to signup")
	}

	return nil
}

func (s *service) SignIn(body *domain.Auth) (string, error) {
	data, err := s.repo.FindByUsername(body)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(body.Password)); err != nil {
		return "", errors.New("incorrect password")
	}

	jwtSecret := "secret"
	claims := JWTClaims{
		ID:       data.ID,
		Username: data.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signed, nil
}
