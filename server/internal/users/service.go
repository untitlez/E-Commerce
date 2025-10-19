package users

import (
	"errors"
)

type service struct {
	repo userRepository
}

func NewService(r userRepository) *service {
	return &service{repo: r}
}

// Get All
func (s *service) getAllUser(filter *query) ([]*User, error) {
	hasFilter := filter.FullName != nil && filter.Email != nil && filter.Password != nil

	if hasFilter {
		search, err := s.repo.findBySearch(filter)
		if err != nil {
			return nil, err
		}

		if len(search) == 0 {
			return nil, errors.New("search not found")
		}

		return search, nil
	}

	if filter.FullName != nil {
		fullName, err := s.repo.findByFullName(filter)
		if err != nil {
			return nil, errors.New("fullname not found")
		}
		return fullName, nil
	}

	if filter.Email != nil {
		email, err := s.repo.findByEmail(filter)
		if err != nil {
			return nil, errors.New("email not found")
		}
		return email, nil
	}

	if filter.Password != nil {
		password, err := s.repo.findByPassword(filter)
		if err != nil {
			return nil, errors.New("password not found")
		}
		return password, nil
	}

	allData, err := s.repo.findAll()
	if err != nil {
		return nil, errors.New("users not found")
	}

	return allData, nil
}

// Get ID
func (s *service) getUser(id int64) (*User, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	data, err := s.repo.findByID(id)
	if err != nil {
		return nil, errors.New("user not found ")
	}

	return data, nil
}

// Create
func (s *service) createUser(body *User) error {
	user := &User{}
	if body == user {
		return errors.New("invalid body")
	}

	if err := s.repo.create(body); err != nil {
		return errors.New("fail to create user")
	}

	return nil
}

// Update
func (s *service) updateUser(id int64, body *User) (*User, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	if err := s.repo.update(id, body); err != nil {
		return nil, errors.New("fail to update user")
	}

	updated, err := s.repo.findByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return updated, nil
}

// Delete
func (s *service) deleteUser(id int64) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	if _, err := s.repo.findByID(id); err != nil {
		return errors.New("user not found")
	}

	if err := s.repo.delete(id); err != nil {
		return errors.New("fail to delete user")
	}

	return nil
}
