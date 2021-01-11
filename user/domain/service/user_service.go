package service

import (
	"errors"
	"go-micro-service/domain/model"
	repository "go-micro-service/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(*model.User) (string, error)
	DeleteUser(string) error
	UpdateUser(user *model.User, isChangedPwd bool) error
	FindUserByName(string) (*model.User, error)
	CheckPwd(userName string, pwd string) (ok bool, err error)
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

// bcrypt
func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//Validation
func ValidatePassword(password string, hash string) (isOK bool, err error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false, errors.New("Wrong password")
	}
	return true, nil
}

func (u *UserDataService) AddUser(user *model.User) (ID string, err error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

func (u *UserDataService) DeleteUser(uid string) error {
	return u.UserRepository.DeleteUserByID(uid)
}

func (u *UserDataService) UpdateUser(user *model.User, isChangedPwd bool) error {
	if isChangedPwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)

	}
	return u.UserRepository.UpdateUser(user)
}

func (u *UserDataService) FindUserByName(name string) (*model.User, error) {
	return u.UserRepository.FindUserByName(name)
}

func (u *UserDataService) CheckPwd(name string, pwd string) (ok bool, err error) {
	user, err := u.UserRepository.FindUserByName(name)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.HashPassword)
}
