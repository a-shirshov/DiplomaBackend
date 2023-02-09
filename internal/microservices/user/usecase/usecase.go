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

func (uU *userUsecase) GetUser(userID int) (*models.User, error) {
	resultUser, err := uU.userRepo.GetUser(userID)
	if err != nil {
		return &models.User{}, err
	}
	resultUser.ImgUrl = utils.TryBuildImgUrl(resultUser.ImgUrl)
	return resultUser, nil
}

func (uU *userUsecase) UpdateUser(user *models.User) (*models.User, error) {
	resultUser, err := uU.userRepo.UpdateUser(user)
	if err != nil {
		return &models.User{}, err
	}

	resultUser.ImgUrl = utils.TryBuildImgUrl(resultUser.ImgUrl)
	return resultUser, nil
}

func (uU *userUsecase) UpdateUserImage(userID int, imgUUID string) (*models.User, error) {
	resultUser, err := uU.userRepo.UpdateUserImage(userID, imgUUID)
	if err != nil {
		return &models.User{}, err
	}

	resultUser.ImgUrl = utils.TryBuildImgUrl(resultUser.ImgUrl)
	return resultUser, nil
}

func (uU *userUsecase) GetFavouriteKudagoEventsIDs(userID int) ([]int, error) {
	return uU.userRepo.GetFavouriteKudagoEventsIDs(userID)
}