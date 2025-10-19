package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo authRepository
}

func NewService(r authRepository) *service {
	return &service{repo: r}
}

func (s *service) signup(body *Auth) error {
	req := body.Username == "" || body.Password == ""
	fmt.Printf("req: %v\n", req)

	if req {
		return errors.New("username or password invalid")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return err
	}

	newBody := &Auth{
		Username: body.Username,
		Password: string(hashPassword),
	}

	if err := s.repo.create(newBody); err != nil {
		return errors.New("fail to signup")
	}

	return nil
}

func (s *service) signin(body *Auth) (string, error) {

	if err := s.repo.findByUsername(body); err != nil {
		return "", errors.New("username already exists")
	}

	data, err := s.repo.findById(int64(body.ID))
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(body.Password)); err != nil {
		return "", errors.New("incorrect password")
	}

	jwtSecret := "secret"

	claims := jwt.MapClaims{
		"username": data.Username,
		"password": data.Password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signed, nil
}
