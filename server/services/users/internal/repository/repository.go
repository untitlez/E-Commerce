package repository

import (
	"server/services/users/internal/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

// Find All
func (r *repository) FindAll() ([]*domain.User, error) {
	users := []*domain.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Find ID
func (r *repository) FindByID(id int64) (*domain.User, error) {
	user := &domain.User{}
	if err := r.db.First(user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Create
func (r *repository) Create(body *domain.User) error {
	return r.db.Create(body).Error
}

// Update
func (r *repository) Update(id int64, body *domain.User) error {
	return r.db.Where("id=?", id).Updates(body).Error
}

// Delete
func (r *repository) Delete(id int64) error {
	return r.db.Unscoped().Delete(&domain.User{}, id).Error
}

// Search
//
// Search FullName
func (r *repository) FindByFullName(filter *domain.Query) ([]*domain.User, error) {
	users := []*domain.User{}
	if err := r.db.Where("full_name ILIKE ?", filter).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Search Email
func (r *repository) FindByEmail(filter *domain.Query) ([]*domain.User, error) {
	users := []*domain.User{}
	if err := r.db.Where("email ILIKE ?", filter).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Search Password
func (r *repository) FindByPassword(filter *domain.Query) ([]*domain.User, error) {
	users := []*domain.User{}
	if err := r.db.Where("password ILIKE ?", filter.Password).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Search All
func (r *repository) FindBySearch(filter *domain.Query) ([]*domain.User, error) {
	users := []*domain.User{}
	if err := r.db.Where("full_name ILIKE ? AND email ILIKE ? AND password ILIKE ?", filter.FullName, filter.Email, filter.Password).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
