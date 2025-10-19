package auth

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) findByUsername(body *Auth) error {
	return r.db.Where("username=?", body.Username).Take(&Auth{}).Error
}

func (r *repository) findById(id int64) (*Auth, error) {
	auth := &Auth{}
	if err := r.db.First(auth, id).Error; err != nil {
		return nil, err
	}

	return auth, nil
}

func (r *repository) create(body *Auth) error {
	return r.db.Create(body).Error
}

func (r *repository) update(body *Auth, id int64) error {
	return r.db.Where("id=?", id).Updates(body).Error
}

func (r *repository) delete(id int64) error {
	return r.db.Unscoped().Delete(&Auth{}, id).Error
}
