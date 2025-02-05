package usecase

import (
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/repository"
)

type UserUseCase interface {
	RegisterNewUser(payload model.UserCredential) (model.UserCredential, error)
	FindAllUser() ([]model.UserCredential, error)
	FindUserById(id uint32) (model.UserCredential, error)
	FindUserByUsernamePassword(username string, password string) (model.UserCredential, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) RegisterNewUser(payload model.UserCredential) (model.UserCredential, error) {
	return u.repo.Create(payload)
}

func (u *userUseCase) FindAllUser() ([]model.UserCredential, error) {
	return u.repo.List()
}

func (u *userUseCase) FindUserById(id uint32) (model.UserCredential, error) {
	return u.repo.Get(id)
}

func (u *userUseCase) FindUserByUsernamePassword(username string, password string) (model.UserCredential, error) {
	// Call the corresponding repository method
	user, err := u.repo.GetByUsernamePassword(username, password)
	if err != nil {
		return model.UserCredential{}, err
	}
	return user, nil
}


func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}

}
