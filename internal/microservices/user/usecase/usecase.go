package Usecase

import (
	"Diploma/internal/microservices/user"
	"Diploma/internal/models"
)

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase(userRepo user.Repository) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (uU *userUsecase) GetUser(userId int) (*models.User, error) {
	return uU.userRepo.GetUser(userId)
}

func (uU *userUsecase) UpdateUser(userId int, user *models.User) (*models.User, error) {
	resultUser, err := uU.userRepo.UpdateUser(userId, user)
	if err != nil {
		return nil, err
	}

	return resultUser, nil
}