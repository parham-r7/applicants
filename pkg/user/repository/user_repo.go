package repository

import (
	"applicants/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	Fetch() (users []*models.User, err error)
	GetByID(userID uint) (user *models.User, err error)
	Create(user *models.User) (usr *models.User, err error)
	Update(user *models.User) (usr *models.User, err error)
	Delete(userID uint) (err error)
}

type userRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB)  UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u userRepo) Fetch() (users []*models.User, err error) {
	panic("implement me")
}

func (u userRepo) GetByID(userID uint) (user *models.User, err error) {
	panic("implement me")
}

func (u userRepo) Create(user *models.User) (usr *models.User, err error) {
	err = u.db.Create(&user).Error
	if err != nil{
		return nil, err
	}

	return usr, nil
}

func (u userRepo) Update(user *models.User) (usr *models.User, err error) {
	panic("implement me")
}

func (u userRepo) Delete(userID uint) (err error) {
	panic("implement me")
}

