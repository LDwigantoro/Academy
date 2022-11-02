package repository

import (
	"latihan3/entity"

	"gorm.io/gorm"
)

// strcut UserRepository
type IUserRepository interface {
	Create(user entity.User) error
	GetAll() ([]entity.User, error)
	Update(user entity.User) error
	Delete(user entity.User) error
	FindById(id int) (entity.User, error)
}

// strcut UserRepository
type UserRepository struct {
	db *gorm.DB
}

// fungsi NewUserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Fungsi CreateUser
func (u UserRepository) Create(user entity.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Fungsi GetListUser
func (u UserRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	if err := u.db.Debug().Find(&users).Error; err != nil {
		return nil, nil
	}
	return users, nil
}

// Fungsi UpdateUser
func (u UserRepository) Update(user entity.User) error {
	if err := u.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// Fungsi DeleteUser
func (u UserRepository) Delete(user entity.User) error {
	if err := u.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

// Fungsi FindById
func (u UserRepository) FindById(id int) (entity.User, error) {
	var user entity.User
	if err := u.db.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
