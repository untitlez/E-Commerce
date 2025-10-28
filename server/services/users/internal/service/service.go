package service

import (
	"errors"
	"server/services/users/internal/domain"
)

type service struct {
	repo domain.UserRepository
}

func NewService(r domain.UserRepository) *service {
	return &service{repo: r}
}

// Get All
func (s *service) GetAllUser(filter *domain.Query) ([]*domain.User, error) {
	hasFilter := filter.FullName != nil && filter.Email != nil && filter.Password != nil

	if hasFilter {
		search, err := s.repo.FindBySearch(filter)
		if err != nil {
			return nil, err
		}

		if len(search) == 0 {
			return nil, errors.New("search not found")
		}

		return search, nil
	}

	if filter.FullName != nil {
		fullName, err := s.repo.FindByFullName(filter)
		if err != nil {
			return nil, errors.New("fullname not found")
		}
		return fullName, nil
	}

	if filter.Email != nil {
		email, err := s.repo.FindByEmail(filter)
		if err != nil {
			return nil, errors.New("email not found")
		}
		return email, nil
	}

	if filter.Password != nil {
		password, err := s.repo.FindByPassword(filter)
		if err != nil {
			return nil, errors.New("password not found")
		}
		return password, nil
	}

	allData, err := s.repo.FindAll()
	if err != nil {
		return nil, errors.New("users not found")
	}

	return allData, nil
}

// Get ID
func (s *service) GetUser(id int64) (*domain.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	data, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found ")
	}

	return data, nil
}

// Create
func (s *service) CreateUser(body *domain.User) error {
	user := &domain.User{}
	if body == user {
		return errors.New("invalid body")
	}

	if err := s.repo.Create(body); err != nil {
		return errors.New("fail to create user")
	}

	return nil
}

// Update
func (s *service) UpdateUser(id int64, body *domain.User) (*domain.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	if err := s.repo.Update(id, body); err != nil {
		return nil, errors.New("fail to update user")
	}

	updated, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return updated, nil
}

// Delete
func (s *service) DeleteUser(id int64) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("user not found")
	}

	if err := s.repo.Delete(id); err != nil {
		return errors.New("fail to delete user")
	}

	return nil
}
