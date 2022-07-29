package Usecase

import (
	"Diploma/internal/microservices/user"
	"Diploma/internal/microservices/auth"
	"Diploma/internal/models"
	"Diploma/utils"
)

type userUsecase struct {
	userRepo user.Repository
	sessionRepo auth.SessionRepository
}

func NewUserUsecase(userRepo user.Repository, sessionRepo auth.SessionRepository) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
		sessionRepo: sessionRepo,
	}
}

func (uU *userUsecase) GetUser(userId int) (*models.User, error) {
	return uU.userRepo.GetUser(userId)
}

func (uU *userUsecase) UpdateUser(au *utils.AccessDetails, user *models.User) (*models.User, error) {
	userId, err := uU.sessionRepo.FetchAuth(au.AccessUuid)
	if err != nil {
		return nil, err
	}

	resultUser, err := uU.userRepo.UpdateUser(userId, user)
	if err != nil {
		return nil, err
	}

	return resultUser, nil
}