package repository

import (
	"SchoolMarket-run-with-go-/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) DeleteUser(Id uint) (int64, error) {
	result := r.db.Delete(&model.User{}, Id)
	return result.RowsAffected, result.Error
}

func (r *UserRepository) FindAim(Id uint) (*model.User, error) {
	var aim_user model.User
	result := r.db.First(&aim_user, Id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &aim_user, nil
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *UserRepository) UpdateUser(user *model.User) (*model.User, error) {
	return user, r.db.Save(user).Error
}
