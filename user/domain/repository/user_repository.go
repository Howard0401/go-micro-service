package repository

import (
	"go-micro-service/domain/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	InitTable() error
	FindUserByName(string) (*model.User, error)
	FindUserByID(string) (*model.User, error)
	CreateUser(*model.User) (string, error)
	DeleteUserByID(string) error
	UpdateUser(*model.User) error
	FindAll() ([]*model.User, error)
}

type UserRepository struct {
	mysqlDB *gorm.DB
}

// input DB
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDB: db}
}

func (u *UserRepository) InitTable() error {
	// return u.mysqlDB.CreateTable(*&model.User{}) //v1
	return u.mysqlDB.Migrator().CreateTable(&model.User{})
}

func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDB.Where("user_name =?", name).Find(user).Error
}

func (u *UserRepository) FindUserByID(uid string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDB.First(user, uid).Error
}

func (u *UserRepository) CreateUser(user *model.User) (id string, err error) {
	return user.ID, u.mysqlDB.Create(user).Error
}

func (u *UserRepository) DeleteUserByID(id string) error {
	return u.mysqlDB.Where("id=?", id).Delete(&model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	// <cf> gorm v1 Update
	return u.mysqlDB.Model(user).Updates(&user).Error
}

func (u *UserRepository) FindAll() (all []*model.User, err error) {
	return all, u.mysqlDB.Find(&all).Error
}
