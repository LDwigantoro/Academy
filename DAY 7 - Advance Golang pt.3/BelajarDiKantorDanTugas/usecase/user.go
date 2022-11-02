package usecase

import (
	"latihan3/entity"
	"latihan3/entity/response"
	"latihan3/repository"

	"github.com/jinzhu/copier"
)

type IUserUsecase interface {
	Create(user entity.User) error
	GetListUser() ([]response.GetUserResponse, error)
	UpdateUser(req response.UpdateUserRequest) error
	DeleteUser(id int) error
}

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepository}
}

func (u UserUsecase) CreateUser(req response.CreateUserRequest) error {

	user := entity.User{}
	copier.Copy(&user, &req)
	err := u.userRepository.Create(user)
	if err != nil {
		return err
	}
	userResponse := []response.GetUserResponse{}
	copier.Copy(&userResponse, &user)
	return nil
}

func (u UserUsecase) GetList() ([]response.GetUserResponse, error) {
	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, nil
	}
	userResponse := []response.GetUserResponse{}
	copier.Copy(&userResponse, &users)
	return userResponse, nil
}

// fungsi UpdateUser
func (u UserUsecase) UpdateUser(req response.UpdateUserRequest) error {
	user := entity.User{}
	copier.Copy(&user, &req)
	err := u.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

// fungsi DeleteUser
func (u UserUsecase) DeleteUser(id int) error {
	_, err := u.userRepository.FindById(id)
	if err != nil {
		return err
	}

	err = u.userRepository.Delete(id)
	return err
}
