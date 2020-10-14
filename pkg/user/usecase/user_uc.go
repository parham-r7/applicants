package usecase

import (
	"applicants/models"
	"applicants/pkg/user/repository"
)

type UserUC interface {
	Fetch() (users []*models.User, err error)
	GetByID(userID uint) (user *models.User, err error)
	Create(user *models.User) (usr *models.User, err error)
	Update(user *models.User) (usr *models.User, err error)
	Delete(userID uint) (err error)
}

type userUC struct {
	repo repository.UserRepo
}

func New(repo repository.UserRepo) UserUC {
	return &userUC{
		repo: repo,
	}
}

func (u userUC) Fetch() (users []*models.User, err error) {
	panic("implement me")
}

func (u userUC) GetByID(userID uint) (user *models.User, err error) {
	panic("implement me")
}

func (u userUC) Create(user *models.User) (usr *models.User, err error) {
	return u.repo.Create(user)
}

func (u userUC) Update(user *models.User) (usr *models.User, err error) {
	panic("implement me")
}

func (u userUC) Delete(userID uint) (err error) {
	panic("implement me")
}


