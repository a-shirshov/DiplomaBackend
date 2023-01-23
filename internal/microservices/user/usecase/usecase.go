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
	resultUser, err :=  uU.userRepo.GetUser(userId)
	if err != nil {
		return &models.User{}, err
	}
	resultUser.ImgUrl = utils.BuildImgUrl(resultUser.ImgUrl)
	return resultUser, nil
}

func (uU *userUsecase) UpdateUser(user *models.User) (*models.User, error) {
	resultUser, err := uU.userRepo.UpdateUser(user)
	if err != nil {
		return &models.User{}, err
	}
	resultUser.ImgUrl = utils.BuildImgUrl(resultUser.ImgUrl)
	return resultUser, nil
}

func (uU *userUsecase) GetFavouriteKudagoEventsIDs(userID int) ([]int, error) {
	favouriteEventIDs, err := uU.userRepo.GetFavouriteKudagoEventsIDs(userID)
	if err != nil {
		return nil, err
	}
	return favouriteEventIDs, nil
}