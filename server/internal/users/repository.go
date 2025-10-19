package users

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

// Find All
func (r *repository) findAll() ([]*User, error) {
	users := []*User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Find ID
func (r *repository) findByID(id int64) (*User, error) {
	user := &User{}
	if err := r.db.First(user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Create
func (r *repository) create(user *User) error {
	return r.db.Create(user).Error
}

// Update
func (r *repository) update(id int64, user *User) error {
	return r.db.Where("id=?", id).Updates(user).Error
}

// Delete
func (r *repository) delete(id int64) error {
	return r.db.Unscoped().Delete(&User{}, id).Error
}

// Search
//
// Search FullName
func (r *repository) findByFullName(filter *query) ([]*User, error) {
	users := []*User{}
	if err := r.db.Where("full_name ILIKE ?", filter).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Search Email
func (r *repository) findByEmail(filter *query) ([]*User, error) {
	users := []*User{}
	if err := r.db.Where("email ILIKE ?", filter).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Search Password
func (r *repository) findByPassword(filter *query) ([]*User, error) {
	users := []*User{}
	if err := r.db.Where("password ILIKE ?", filter.Password).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Search All
func (r *repository) findBySearch(filter *query) ([]*User, error) {
	users := []*User{}
	if err := r.db.Where("full_name ILIKE ? AND email ILIKE ? AND password ILIKE ?", filter.FullName, filter.Email, filter.Password).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
