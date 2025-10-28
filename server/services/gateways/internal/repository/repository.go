package repository

import (
	"server/services/gateways/internal/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindByUsername(body *domain.Auth) (*domain.Auth, error) {
	auth := &domain.Auth{}
	if err := r.db.Where("username=?", body.Username).Take(auth).Error; err != nil {
		return nil, err
	}

	return auth, nil
}

func (r *repository) FindById(id int64) (*domain.Auth, error) {
	auth := &domain.Auth{}
	if err := r.db.First(auth, id).Error; err != nil {
		return nil, err
	}

	return auth, nil
}

func (r *repository) Create(body *domain.Auth) error {
	return r.db.Create(body).Error
}

func (r *repository) Update(body *domain.Auth, id int64) error {
	return r.db.Where("id=?", id).Updates(body).Error
}

func (r *repository) Delete(id int64) error {
	return r.db.Unscoped().Delete(&domain.Auth{}, id).Error
}
