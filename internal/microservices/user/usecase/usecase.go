package usecase

import (
	"Diploma/internal/microservices/user"
	"Diploma/internal/models"
	"Diploma/utils"
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

func (uU *userUsecase) UpdateUser(user *models.User) (*models.User, error) {
	resultUser, err := uU.userRepo.UpdateUser(user)
	if err != nil {
		return &models.User{}, err
	}
	resultUser.ImgUrl = utils.BuildImgUrl(resultUser.ImgUrl)
	return resultUser, nil
}